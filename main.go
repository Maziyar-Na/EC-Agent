// Btw, the reason this has been written as one script as of now is because the code in main() will be inserted into
// the kubelet and the rest of the file will be a standalone package

package main

import (
	"bufio"
	"context"
	"encoding/binary"
	"fmt"
	pb "github.com/Maziyar-Na/EC-Agent/grpc"
	msg_struct "github.com/Maziyar-Na/EC-Agent/msg"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

const MAXGCMNO = 30
const PORT = ":4445"
const PORT_GRPC	= ":4446"
const BUFFSIZE = 2048
const EC_CONNECT_SYSCALL = 335
const RESIZE_MEM_SYSCALL = 336
const INCREASE_MEM_CG_MARGIN_SYSCALL = 337
const RESIZE_QUOTA_SYSCALL = 338
const READ_QUOTA_SYSCALL = 339
const GET_PARENT_CGID_SYSCALL = 340

const INTERFACE = "eno1" // This could be changed
//const INTERFACE = "enp0s3"

type server struct {
	pb.UnimplementedHandlerServer
}

func getIpFromInterface(inter string) net.IP {
	byNameInterface, err := net.InterfaceByName(inter)
	if err != nil {
		log.Println(err)
	}
	addresses, err := byNameInterface.Addrs()
	// for k, v := range addresses {
	// 	log.Printf("Interface Address #%v : %v\n", k, v.String())
	// }
	ipv4Addr, _, err := net.ParseCIDR(addresses[0].String())

	return ipv4Addr
}

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

func GetDockerPid(dockerId string) (int, int, string) {
	cmd := "sudo docker inspect --format '{{ .State.Pid }}' " + dockerId
	out, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return 0, 1, "ERROR in getting PID of container with id " + dockerId + ": " + err.Error()
	}
	pid := string(out)
	pid = strings.TrimSuffix(pid, "\n")
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		return  0, 1, err.Error()
	}
	return pidInt, 0, ""
}

func RunConnectContainer(gcmIpStr string, dockerId string, pid int) (string, int32, uint64){
	// call syscall for ec_connect here
	gcmIp := ip2int(net.ParseIP(gcmIpStr))
	port := 4444
	interfaceIP := getIpFromInterface(INTERFACE)
	//log.Printf("[INFO]: IP of the interface %s is %s\n", INTERFACE, interfaceIP)
	//agentIP := ip2int(net.ParseIP("128.105.144.93"))
	agentIP := ip2int(interfaceIP)
	cgId, t, err := syscall.Syscall6(EC_CONNECT_SYSCALL, uintptr(gcmIp) , uintptr(port), uintptr(pid), uintptr(agentIP) , 0, 0)

	log.Println("cgID: " + string(int(cgId)) + ", t: " + string(t) + ", err: " + err.Error())
	if int(cgId) == -1 {
		fmt.Println("ERROR IN RUNCONNECT CONTAINER. Rx back cgroupID: -1")
	}
	return dockerId, int32(cgId), 0
}

// ReqContainerInfo implements agent.HandlerServer
func (s *server) ReqConnectContainer(ctx context.Context, in *pb.ConnectContainerRequest) (*pb.ConnectContainerReply, error) {
	log.Printf("Received: %v, %v, %v", in.GetGcmIP(), in.GetPodName(), in.GetDockerId())
	pid, ret, err := GetDockerPid(in.GetDockerId())
	if ret != 0 {
		log.Println("Error getting docker pid for container: " + in.GetDockerId() + ", Err: " + err)
		log.Println()
	}
	_, cgroupId, val := RunConnectContainer(in.GcmIP, in.GetDockerId(), pid)
	if val != 0 {
		log.Println("Error getting docker pid for container: " + in.GetDockerId() + ", Err: " + string(val))
		log.Println()
	}
	fmt.Print(pid)
	if int(cgroupId) == -1 {
		fmt.Println("ERROR IN REQCONNECT CONTAINER. Rx back cgroupID: -1")
	}
	//var cgroupId int32
	//cgroupId = 42

	return &pb.ConnectContainerReply{
		PodName: in.GetPodName(),
		DockerID: in.GetDockerId(),
		CgroupID: cgroupId,
	}, nil
}

func handleCpuReq(cgroupId int32, quota uint64) (uint64, uint64) {
	//log.Printf("setting quota to: %d\n", quota)
	var updatedQuota uint64
	quotaMega := quota/1000
	var fistCgroupToUpdate int32
	var secondCgroupToUpdate int32
	var isInc int32

	//Getting Current quota -> compare with new quota -> conclude which cgroup to update first
	curr_quota, _, _ := syscall.Syscall(READ_QUOTA_SYSCALL, uintptr(cgroupId), 0, 0)
	currQuota := uint64(curr_quota)
	//log.Println("[INFO] cuurent quota: ", currQuota)
	//log.Println("[INFO] new quota:", quotaMega)
	if currQuota < quotaMega {
		isInc = 1
	} else {
		//log.Println("[INFO] we are decreasing the quota!")
		isInc = 0
	}

	parentCgroupID, _, _ := syscall.Syscall(GET_PARENT_CGID_SYSCALL, uintptr(cgroupId), 0, 0)
	parentCgID := int32(parentCgroupID)
	//log.Println("getting the parent id: ", int32(parentCgID))
	//TODO: need error handling here
	if isInc == 1 {
		fistCgroupToUpdate = parentCgID
		secondCgroupToUpdate = cgroupId
	} else {
		fistCgroupToUpdate = cgroupId
		secondCgroupToUpdate = parentCgID
	}

	//Which cgroup to update first is clear now -> let's do it
	ret, _, _ := syscall.Syscall(RESIZE_QUOTA_SYSCALL, uintptr(fistCgroupToUpdate), uintptr(quotaMega), 0)
	if ret == 1 {
		log.Println("[Error] Quota Set Failed at the first level!")
		ret = 1
		updatedQuota = 0
		return updatedQuota, uint64(ret)
	}

	ret, _, _ = syscall.Syscall(RESIZE_QUOTA_SYSCALL, uintptr(secondCgroupToUpdate), uintptr(quotaMega), 0)
	if ret == 1 {
		log.Println("Quota Set Failed at the second level!")
		ret = 1
		updatedQuota = 0
	} else {
		//log.Println("Quota Set Success. set to: ", uint64(ret))
		updatedQuota = uint64(ret)
		ret = 0
	}

	return updatedQuota, uint64(ret)
}

func handleMemReq(cgroupId int32) uint64 {
	log.Printf("cgroup_id: %d\n", cgroupId)
	availMemRet, _, _ := syscall.Syscall(RESIZE_MEM_SYSCALL, uintptr(cgroupId), 0, 0)
	availMem := uint64(availMemRet)

	log.Printf("[INFO]: EC Agent: Reclaimed memory is: %d\n", availMem)
	return availMem
}
//Assumption: we deploy a single container per pod, when we want to resize,
//first we change the memory limit of the pod then the target container itself
func handleResizeMaxMem(cgroupId int32, newLimit uint64, isMemsw int, isInc int) uint64 {
	log.Printf("setting new mem limit to: %d\n", newLimit)
	var fistCgroupToUpdate int32
	var secondCgroupToUpdate int32


	parentCgroupID, _, _ := syscall.Syscall(GET_PARENT_CGID_SYSCALL, uintptr(cgroupId), 0, 0)
	parentCgID := int32(parentCgroupID)
	
	if isInc == 1 {
		fistCgroupToUpdate = parentCgID
		secondCgroupToUpdate = cgroupId
	} else {
		fistCgroupToUpdate = cgroupId
		secondCgroupToUpdate = parentCgID
	}
	//TODO: error handling needed here
	availMemRet, a, b := syscall.Syscall(RESIZE_MEM_SYSCALL, uintptr(fistCgroupToUpdate), uintptr(newLimit), uintptr(isMemsw))
	availMem := uint64(availMemRet)
	fmt.Println(int64(a))
	fmt.Println(int64(b))

	if availMem != 0 {
		log.Printf("[INFO]: EC Agent: resize_max_mem fails in first level. Ret: %d \n", availMem)
		return availMem
	}

	availMemRet, _, _ = syscall.Syscall(RESIZE_MEM_SYSCALL, uintptr(secondCgroupToUpdate), uintptr(newLimit), uintptr(isMemsw))
	availMem = uint64(availMemRet)

	if availMem != 0 {
		log.Printf("[INFO]: EC Agent: resize_max_mem fails in second level. Ret: %d \n", availMem)
	}

	return availMem
}

func handleConnection(conn net.Conn) {
	log.Printf("[DBG] Server: New fd created for new connection. Serving %s\n", conn.RemoteAddr().String())
	for {
		buff := make([]byte, BUFFSIZE)
		c := bufio.NewReader(conn)
		defer conn.Close()
		// read a single byte which contains the message length at the beginning of the message
		size, err := c.ReadByte()
		if err != nil {
			if err.Error() == "EOF" {
				log.Println("Connection killed by client")
				break
			} else {
				log.Println("ERROR in reading Header: ", err.Error())
			}
		}
		//log.Println("[ProtoBuf] RX Message Body length: ", size)
		// now, read the full Protobuf message
		_, err = io.ReadFull(c, buff[:int(size)])
		if err != nil {
			log.Println("ERROR in reading Body: ", err.Error())
		}
		rxMsg := &msg_struct.ECMessage{}
		err = proto.Unmarshal(buff[:size], rxMsg)
		if err != nil {
			log.Println("ERROR in ProtoBuff - UnMarshaling: ", err.Error())
		}

		// log.Println("Recieved message req type: ", rxMsg.GetReqType())
		var ret uint64
		var container_id string
		var updated_quota uint64
		//log.Println("--------------- BEGIN NEW REQUEST ---------------")
		switch rxMsg.GetReqType() {
		case 0:
			//log.Println("CPU Request")
			updated_quota, ret = handleCpuReq(rxMsg.GetCgroupId(), rxMsg.GetQuota())
		case 1:
			//log.Println("Memory Request")
			ret = handleMemReq(rxMsg.GetCgroupId())
		case 2:
			log.Println("Init Request")
		case 3:
			log.Println("CPU SLICE")
		case 5:
			//log.Println("Handle RESIZE MAX/MIN")
			ret = handleResizeMaxMem(rxMsg.GetCgroupId(), rxMsg.GetRsrcAmnt(), 0, 0)
		default:
			log.Println("[ERROR] Not going in the right way! request type is invalid!")
		}
		//log.Println("--------------- END NEW REQUEST ---------------")

		//log.Println("Docker Container id:", container_id)
		//log.Println("Updated Quota", updated_quota)
		txMsg := &msg_struct.ECMessage{
			ReqType: rxMsg.GetReqType(),
			RsrcAmnt: ret,
			Quota: updated_quota,
			PayloadString: container_id,
			Request: rxMsg.GetRequest(),
		}

		txMsgMarshal, err := proto.Marshal(txMsg)
		if err != nil {
			log.Fatal("TX Data marshaling error: ", err)
		}

		// Write to socket the message stream
		_, err = conn.Write(txMsgMarshal)
		if err != nil {
			log.Println("[ERROR] in writing proto message to socket" + err.Error())
		}
		//log.Printf("[PROTOBUF] TX  Message Body length: %d\n", length)

	}
}


func GrpcServer(wg *sync.WaitGroup) {
	defer wg.Done()
	l, err := net.Listen("tcp4", PORT_GRPC)
	if err != nil {
		log.Println(err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterHandlerServer(s, &server{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func TcpServer(wg *sync.WaitGroup) {
	defer wg.Done()
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Listening on port: " + PORT)
	for {
		if conn, err := l.Accept(); err == nil {
			go handleConnection(conn)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go GrpcServer(&wg)
	go TcpServer(&wg)
	wg.Wait()
}
