// Btw, the reason this has been written as one script as of now is because the code in main() will be inserted into
// the kubelet and the rest of the file will be a standalone package

package main

import (
	"context"
	"encoding/binary"
	pb "github.com/Maziyar-Na/EC-Agent/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const MAXGCMNO = 30
const PORT = ":4445"
const BUFFSIZE = 2048
const EC_CONNECT_SYSCALL = 335
const RESIZE_MEM_SYSCALL = 336
const INCREASE_MEM_CG_MARGIN_SYSCALL = 337
const RESIZE_QUOTA_SYSCALL = 338
const READ_QUOTA_SYSCALL = 339

//const INTERFACE = "eno1" // This could be changed
const INTERFACE = "enp0s3"

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

func connectContainer(serverIp, containerName string) (string, int32, uint64) {
	log.Printf("[DBG] CONNECT CONTAINER: server_ip %s: , container name: %s\n", serverIp, containerName)
	cmdForDockId := "sudo docker ps -a | grep k8s_" + containerName + " | awk '{print $1}'"
	var containerId string
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	for {
		out, err := exec.CommandContext(ctx, "/bin/sh",  "-c" , cmdForDockId).Output()
		//println("container id: ", string(out))
		if len(string(out)) > 0 {
			containerId = string(out)
			break
		}
		if ctx.Err() != nil {
			return "Error in finding container: " + err.Error(), 0, 1
		}
	}
	containerId = strings.TrimSuffix(containerId, "\n")
	log.Println("got container id!: " + containerId)
	cmd := "sudo docker inspect --format '{{ .State.Pid }}' " + containerId
	out, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		return "ERROR in getting PID of container with id " + containerId + ": " + err.Error(), 0, 1
	}
	pid := string(out)
	pid = strings.TrimSuffix(pid, "\n")
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		return err.Error(), 0, 1
	}

	log.Println("got pid: " + pid)

	// call syscall for ec_connect here
	gcmIp := ip2int(net.ParseIP(serverIp))
	port := 4444
	interfaceIP := getIpFromInterface(INTERFACE)
	log.Printf("[INFO]: IP of the interface %s is %s\n", INTERFACE, interfaceIP)
	//agentIP := ip2int(net.ParseIP("128.105.144.93"))
	agentIP := ip2int(interfaceIP)
	cgId, _, err := syscall.Syscall6(EC_CONNECT_SYSCALL, uintptr(gcmIp) , uintptr(port), uintptr(pidInt), uintptr(agentIP) , 0, 0)

	log.Println("Docker Container id: " + containerId + ", cgroupId: " + string(cgId))
	return containerId, int32(cgId), 0
}



//func ConnectContainerGrpc(clientIP, )

// ReqContainerInfo implements agent.HandlerServer
func (s *server) ReqContainerInfo(ctx context.Context, in *pb.ContainerRequest) (*pb.ContainerReply, error) {
	log.Printf("Received: %v, %v", in.GetGcmIP(), in.GetPodName())
	containerId, cgId, ret := connectContainer(in.GetGcmIP(), in.GetPodName())
	if ret != 0 {
		log.Println("[ERROR] Initial Container Connection failed...")
		log.Println(containerId)
	}
	return &pb.ContainerReply{
		PodName: in.GetPodName(),
		DockerID: containerId,
		CgroupID: cgId,
	}, nil

	//return &pb.ContainerReply{
	//	ClientIP: "YO:" + in.GetClientIP(),
	//	PodName:"POD_NAME: " + in.GetPodName(),
	//	DockerID: "DOCKER_ID",
	//	CgroupID: 23,
	//}, nil
}

func main() {
	l, err := net.Listen("tcp4", PORT)
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