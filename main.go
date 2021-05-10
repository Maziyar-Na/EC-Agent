// Btw, the reason this has been written as one script as of now is because the code in main() will be inserted into
// the kubelet and the rest of the file will be a standalone package

package main

import (
	"context"
	"encoding/binary"
	"fmt"
	pbController "github.com/Maziyar-Na/EC-Agent/containerUpdateGrpc"
	pbDeployer "github.com/Maziyar-Na/EC-Agent/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

//const TcpPort = ":4445"
const PortGrpcDeployer = ":4446"
const PortGrpcController =":4448"
const BuffSize = 2048
const EcConnectSyscall = 335
const ResizeMemSyscall = 336
const ResizeQuotaSyscall = 338
const GetParentCgidSyscall = 340
const BaseTcpPort = 5000
const BaseUdpPort = 6000
const ReadMemUsageSyscall = 341
const ReadMemLimitSyscall = 342

const INTERFACE = "enp0s3"

type grpcDeployerServer struct {
	pbDeployer.UnimplementedHandlerServer
}

type grpcControllerServer struct {
	pbController.UnimplementedContainerUpdateHandlerServer
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
	fmt.Println("Get docker PID returned pid: " + pid)
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		return  0, 1, err.Error()
	}
	return pidInt, 0, ""
}

func RunConnectContainer(gcmIpStr string, dockerId string, pid int, appNum int32) (string, int32, uint64){
	// call syscall for ec_connect here
	gcmIp := ip2int(net.ParseIP(gcmIpStr))
	//port_tcp := 4444
	//port_udp := 4447
	port_tcp := BaseTcpPort + appNum - 1
	port_udp := BaseUdpPort + appNum - 1
	interfaceIP := getIpFromInterface(INTERFACE)
	//log.Printf("[INFO]: IP of the interface %s is %s\n", INTERFACE, interfaceIP)
	//agentIP := ip2int(net.ParseIP("128.105.144.93"))
	fmt.Println("pid to run connectContainer: " + strconv.Itoa(pid))
	agentIP := ip2int(interfaceIP)
	cgId, t, err := syscall.Syscall6(EcConnectSyscall, uintptr(gcmIp) , uintptr(port_tcp), uintptr(port_udp), uintptr(pid), uintptr(agentIP), 0)

	log.Println("cgID: " + strconv.Itoa(int(cgId)) + ", t: " + string(t) + ", err: " + err.Error())
	if int(cgId) == -1 {
		fmt.Println("ERROR IN RUNCONNECT CONTAINER. Rx back cgroupID: -1")
	}
	return dockerId, int32(cgId), 0
}

// ReqContainerInfo implements agent.HandlerServer
func (s *grpcDeployerServer) ReqConnectContainer(ctx context.Context, in *pbDeployer.ConnectContainerRequest) (*pbDeployer.ConnectContainerReply, error) {
	log.Printf("Received: %v, %v, %v, %d", in.GetGcmIP(), in.GetPodName(), in.GetDockerId(), in.GetAppNum())
	pid, ret, err := GetDockerPid(in.GetDockerId())
	cgroupId := int32(0)
	if ret != 0 {
		log.Println("Error getting docker pid for container: " + in.GetDockerId() + ", Err: " + err)
		log.Println()
	} else {
		_, cgId, val := RunConnectContainer(in.GcmIP, in.GetDockerId(), pid, in.GetAppNum())
		if val != 0 {
			log.Println("Error getting docker pid for container: " + in.GetDockerId() + ", Err: " + string(val))
			log.Println()
		} else {
			cgroupId = cgId
		}
		fmt.Print(pid)
		if int(cgroupId) == -1 {
			fmt.Println("ERROR IN REQCONNECT CONTAINER. Rx back cgroupID: -1")
		}
	}

	return &pbDeployer.ConnectContainerReply{
		PodName: in.GetPodName(),
		DockerID: in.GetDockerId(),
		CgroupID: cgroupId,
	}, nil
}

func (s *grpcControllerServer) ReqQuotaUpdate(ctx context.Context, in *pbController.ContainerQuotaRequest) (*pbController.ContainerQuotaReply, error) {
	//log.Printf("Received: %v, %v, %v, %d", in.GetCgroupId(), in.GetNewQuota(), in.GetResizeFlag(), in.GetSequenceNum())
	var updatedQuota uint64
	quotaMega := in.GetNewQuota()/1000
	change := in.GetResizeFlag()
	var fistCgroupToUpdate int32
	var secondCgroupToUpdate int32

	parentCgroupID, _, _ := syscall.Syscall(GetParentCgidSyscall, uintptr(in.GetCgroupId()), 0, 0)
	parentCgID := int32(parentCgroupID)
	//log.Println("getting the parent id: ", int32(parentCgID))
	//TODO: need error handling here
	if change == "incr" {
		fistCgroupToUpdate = parentCgID
		secondCgroupToUpdate = in.GetCgroupId()
	} else if change == "decr" {
		fistCgroupToUpdate = in.GetCgroupId()
		secondCgroupToUpdate = parentCgID
	} else {
		log.Println("[Error] change is neither incr or decr! it is: " + change)
	}

	//Which cgroup to update first is clear now -> let's do it
	ret, _, _ := syscall.Syscall(ResizeQuotaSyscall, uintptr(fistCgroupToUpdate), uintptr(quotaMega), 0)
	if ret == 1 {
		log.Println("[Error] Quota Set Failed at the first level!")
		ret = 1
		updatedQuota = 0
		//return updatedQuota, uint64(ret)
		return &pbController.ContainerQuotaReply{
			CgroupId: in.GetCgroupId(),
			UpdateQuota: updatedQuota,
			ErrorCode: int32(ret),
			SequenceNum: in.GetSequenceNum(),
		}, nil
	}

	ret, _, _ = syscall.Syscall(ResizeQuotaSyscall, uintptr(secondCgroupToUpdate), uintptr(quotaMega), 0)
	if ret == 1 {
		log.Println("Quota Set Failed at the second level!")
		ret = 1
		updatedQuota = 0
	} else {
		//log.Println("Quota Set Success. set to: ", uint64(ret))
		updatedQuota = uint64(ret)
		ret = 0
	}

	return &pbController.ContainerQuotaReply{
		CgroupId: in.GetCgroupId(),
		UpdateQuota: updatedQuota,
		ErrorCode: int32(ret),
		SequenceNum: in.GetSequenceNum(),
	}, nil

}

func (s *grpcControllerServer) ReqResizeMaxMem(ctx context.Context, in *pbController.ResizeMaxMemRequest) (*pbController.ResizeMaxMemReply, error) {

	cgroupId := in.CgroupId
	newLimit := in.NewMemLimit
	isMemsw := 0
	//isInc := 0

	parentCgroupID, _, _ := syscall.Syscall(GetParentCgidSyscall, uintptr(cgroupId), 0, 0)
	parentCgID := int32(parentCgroupID)

	//always decr based on setup
	fistCgroupToUpdate := cgroupId
	secondCgroupToUpdate := parentCgID

	//if isInc == 1 {
	//	fistCgroupToUpdate = parentCgID
	//	secondCgroupToUpdate = cgroupId
	//} else {
	//	fistCgroupToUpdate = cgroupId
	//	secondCgroupToUpdate = parentCgID
	//}
	//TODO: error handling needed here

	errVal, _, _ := syscall.Syscall(ResizeMemSyscall, uintptr(fistCgroupToUpdate), uintptr(newLimit), uintptr(isMemsw))
	err := uint64(errVal)

	if err != 0 {
		log.Printf("[INFO]: EC Agent: resize_max_mem fails in first level. Ret: %d \n", err)
		//return err
		return &pbController.ResizeMaxMemReply{
			CgroupId:  in.GetCgroupId(),
			ErrorCode: int32(err),
		}, nil
	}

	errVal, _, _ = syscall.Syscall(ResizeMemSyscall, uintptr(secondCgroupToUpdate), uintptr(newLimit), uintptr(isMemsw))
	err = uint64(errVal)

	if err != 0 {
		log.Printf("[INFO]: EC Agent: resize_max_mem fails in second level. Ret: %d \n", err)
	} else {
		log.Printf("Successfuly resized mem for cgid %d to: %d\n", cgroupId, newLimit)
	}

	return &pbController.ResizeMaxMemReply{
		CgroupId:  in.GetCgroupId(),
		ErrorCode: int32(err),
	}, nil

	//return err //err should be 0!

}

func (s *grpcControllerServer) ReadMemUsage(ctx context.Context, in *pbController.CgroupId) (*pbController.ReadMemUsageReply, error) {
	memUsageRet, _, _ := syscall.Syscall(ReadMemUsageSyscall, uintptr(in.GetCgroupId()), 0, 0)
	memUsage := uint64(memUsageRet)

	log.Printf("[INFO]: Memory usage for cgid: %d is: %d\n", in.GetCgroupId(), memUsage)
	return &pbController.ReadMemUsageReply{
		CgroupId: in.GetCgroupId(),
		MemUsage: int64(memUsage),
	}, nil

}

func (s *grpcControllerServer) ReadMemLimit(ctx context.Context, in *pbController.CgroupId) (*pbController.ReadMemLimitReply, error) {
	memLimitRet, _, _ := syscall.Syscall(ReadMemLimitSyscall, uintptr(in.GetCgroupId()), 0, 0)
	memLimit := uint64(memLimitRet)

	log.Printf("[INFO]: Memory Limit for cgid: %d is: %d\n", in.GetCgroupId(), memLimit)
	return &pbController.ReadMemLimitReply{
		CgroupId: in.GetCgroupId(),
		MemLimit: int64(memLimit),
	}, nil

}

//func handleConnection(conn net.Conn) {
//	log.Printf("[DBG] Server: New fd created for new connection. Serving %s\n", conn.RemoteAddr().String())
//	for {
//		buff := make([]byte, BuffSize)
//		c := bufio.NewReader(conn)
//		defer conn.Close()
//		// read a single byte which contains the message length at the beginning of the message
//		size, err := c.ReadByte()
//		if err != nil {
//			if err.Error() == "EOF" {
//				log.Println("Connection killed by client")
//				break
//			} else {
//				log.Println("ERROR in reading Header: ", err.Error())
//			}
//		}
//		//log.Println("[ProtoBuf] RX Message Body length: ", size)
//		// now, read the full Protobuf message
//		_, err = io.ReadFull(c, buff[:int(size)])
//		if err != nil {
//			log.Println("ERROR in reading Body: ", err.Error())
//		}
//		rxMsg := &msg_struct.ECMessage{}
//		err = proto.Unmarshal(buff[:size], rxMsg)
//		if err != nil {
//			log.Println("ERROR in ProtoBuff - UnMarshaling: ", err.Error())
//		}
//
//		// log.Println("Recieved message req type: ", rxMsg.GetReqType())
//		var ret uint64
//		var container_id string
//		var updated_quota uint64
//		//log.Println("--------------- BEGIN NEW REQUEST ---------------")
//		switch rxMsg.GetReqType() {
//		case 0:
//			//log.Println("CPU Request")
//			updated_quota, ret = handleCpuReq(rxMsg.GetCgroupId(), rxMsg.GetQuota(), rxMsg.GetPayloadString())
//		case 1:
//			//log.Println("Memory Request")
//			ret = handleMemReq(rxMsg.GetCgroupId())
//		case 2:
//			log.Println("Init Request")
//		case 3:
//			log.Println("CPU SLICE")
//		case 5:
//			//log.Println("Handle RESIZE MAX/MIN")
//			ret = handleResizeMaxMem(rxMsg.GetCgroupId(), rxMsg.GetRsrcAmnt(), 0, 0)
//		case 6:
//			ret = readMemUsage(rxMsg.GetCgroupId())
//		case 7:
//			ret = readMemLimit(rxMsg.GetCgroupId())
//		default:
//			log.Println("[ERROR] Not going in the right way! request type is invalid!")
//		}
//		//log.Println("--------------- END NEW REQUEST ---------------")
//
//		//log.Println("Docker Container id:", container_id)
//		//log.Println("Updated Quota", updated_quota)
//		txMsg := &msg_struct.ECMessage{
//			ReqType: rxMsg.GetReqType(),
//			RsrcAmnt: ret,
//			Quota: updated_quota,
//			PayloadString: container_id,
//			Request: rxMsg.GetRequest(),
//		}
//
//		txMsgMarshal, err := proto.Marshal(txMsg)
//		if err != nil {
//			log.Fatal("TX Data marshaling error: ", err)
//		}
//
//		// Write to socket the message stream
//		_, err = conn.Write(txMsgMarshal)
//		if err != nil {
//			log.Println("[ERROR] in writing proto message to socket" + err.Error())
//		}
//		//log.Printf("[PROTOBUF] TX  Message Body length: %d\n", length)
//
//	}
//}


func GrpcServerDeployer(wg *sync.WaitGroup) {
	defer wg.Done()
	l, err := net.Listen("tcp4", PortGrpcDeployer)
	if err != nil {
		log.Println(err)
		return
	}
	s := grpc.NewServer()
	log.Println("Grpc Deployer Listening on port: " + PortGrpcDeployer)
	pbDeployer.RegisterHandlerServer(s, &grpcDeployerServer{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func GrpcServerController(wg *sync.WaitGroup) {
	defer wg.Done()
	l, err := net.Listen("tcp4", PortGrpcController)
	if err != nil {
		log.Println(err)
		return
	}
	s := grpc.NewServer()
	log.Println("Grpc Controller Listening on port: " + PortGrpcController)
	pbController.RegisterContainerUpdateHandlerServer(s, &grpcControllerServer{})
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

//func TcpServer(wg *sync.WaitGroup) {
//	defer wg.Done()
//	l, err := net.Listen("tcp4", TcpPort)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	log.Println("Listening on port: " + TcpPort)
//	for {
//		if conn, err := l.Accept(); err == nil {
//			go handleConnection(conn)
//		}
//	}
//}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go GrpcServerDeployer(&wg)
	go GrpcServerController(&wg)
	//go TcpServer(&wg)
	wg.Wait()
}
