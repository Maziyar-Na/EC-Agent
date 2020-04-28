// Btw, the reason this has been written as one script as of now is because the code in main() will be inserted into
// the kubelet and the rest of the file will be a standalone package

package main

import (
	"context"
	"encoding/binary"
	"flag"
	pb "github.com/Maziyar-Na/EC-Agent/grpc"
	"google.golang.org/grpc"
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"net"
	"os/exec"
	"path/filepath"
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

var clientset *kubernetes.Clientset

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

func GetPodFromName(podName string) *v1.Pod {
	podObj, _ := clientset.CoreV1().Pods(apiv1.NamespaceDefault).Get(context.TODO(), podName, metav1.GetOptions{})
	return podObj
}

func GetDockerId(pod *v1.Pod) string {
	return pod.Status.ContainerStatuses[0].ContainerID[9:]
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
	log.Printf("[INFO]: IP of the interface %s is %s\n", INTERFACE, interfaceIP)
	//agentIP := ip2int(net.ParseIP("128.105.144.93"))
	agentIP := ip2int(interfaceIP)
	cgId, t, err := syscall.Syscall6(EC_CONNECT_SYSCALL, uintptr(gcmIp) , uintptr(port), uintptr(pid), uintptr(agentIP) , 0, 0)

	log.Println("cgID: " + string(cgId) + ", t: " + string(t) + ", err: " + err.Error())
	return dockerId, int32(cgId), 0
}

//func ConnectContainerGrpc(clientIP, )

// ReqContainerInfo implements agent.HandlerServer
func (s *server) ReqContainerInfo(ctx context.Context, in *pb.ContainerRequest) (*pb.ContainerReply, error) {
	log.Printf("Received: %v, %v", in.GetGcmIP(), in.GetPodName())
	pod := GetPodFromName(in.GetPodName())
	dockerId := GetDockerId(pod)
	pid, ret, err := GetDockerPid(dockerId)
	if ret != 0 {
		log.Println("Error getting docker pid for container: " + dockerId + ", Err: " + err)
		log.Println()
	}
	_, cgroupId, val := RunConnectContainer(in.GcmIP, dockerId, pid)
	if val != 0 {
		log.Println("Error getting docker pid for container: " + dockerId + ", Err: " + string(val))
		log.Println()
	}

	return &pb.ContainerReply{
		PodName: in.GetPodName(),
		DockerID: dockerId,
		CgroupID: cgroupId,
	}, nil
}

func main() {
	clientset = configK8()

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

func configK8() *kubernetes.Clientset {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return clientset
}