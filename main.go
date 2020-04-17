// Btw, the reason this has been written as one script as of now is because the code in main() will be inserted into 
// the kubelet and the rest of the file will be a standalone package

package main

import (
	"log"
	"net"
	"bufio"
	"io"
	"github.com/golang/protobuf/proto"
	"msg_struct"
	"os/exec"
	"time"
	"context"
	"syscall"
	"strconv"
	"strings"
	"encoding/binary"
)

const MAXGCMNO = 30
const PORT = ":4445"
const BUFFSIZE = 2048
const EC_CONNECT_SYSCALL = 335
const RESIZE_MEM_SYSCALL = 336
const INCREASE_MEM_CG_MARGIN_SYSCALL = 337
const RESIZE_QUOTA_SYSCALL = 338
const READ_QUOTA_SYSCALL = 339

const INTERFACE = "eno1" // This could be changed

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

func connect_container(server_ip, container_name string) (string, uint64) {
	log.Printf("[DBG] CONNECT CONTAINER: server_ip %s: , container name: %s\n", server_ip, container_name)
	
	// Heads up, this will change once it's integrated into kubelet (cuz kubelets check for alive pods/containers in a different way)
	cmd := "sudo docker ps -a | grep k8s_" + container_name + " | awk '{print $1}'"
	// Loops for max of 10 seconds or whenever uit finds the container, whichever comes first
	var container_id string
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for {
    	out, err := exec.CommandContext(ctx, "/bin/sh",  "-c" , cmd).Output()
		if len(string(out)) > 0 {
			container_id = string(out)
			break
		}
		if (ctx.Err() != nil) {
			return "Error in finding container: " + err.Error(), 1
		}
	}
	container_id = strings.TrimSuffix(container_id, "\n")

	cmd = "sudo docker inspect --format '{{ .State.Pid }}' " + container_id
	out, err := exec.Command("/bin/sh", "-c", cmd).Output()  
	if err != nil {
		return "ERROR in getting PID of container with id " + container_id + ": " + err.Error(), 1
	}
	pid := string(out)
	pid = strings.TrimSuffix(pid, "\n")
	pid_int, err := strconv.Atoi(pid)
    if err != nil {
		return err.Error(), 1
	}

	// call syscall for ec_connect here
	gcm_ip := ip2int(net.ParseIP(server_ip))
	port := 4444
	interfaceIP := getIpFromInterface(INTERFACE)
	log.Printf("[INFO]: IP of the interface %s is %s\n", INTERFACE, interfaceIP)
	//agentIP := ip2int(net.ParseIP("128.105.144.93"))
	agentIP := ip2int(interfaceIP)
	_, _, err = syscall.Syscall6(EC_CONNECT_SYSCALL, uintptr(gcm_ip) , uintptr(port), uintptr(pid_int), uintptr(agentIP) , 0, 0)
	
	// log.Println("Docker Container id:", container_id)
	return container_id, 0
}

func handle_cpu_req(cgroup_id, quota uint64) (uint64, uint64) {
	log.Printf("setting quota to: %d\n", quota)
	quotaMega := quota/1000
	update_quota, _, _ := syscall.Syscall(RESIZE_QUOTA_SYSCALL, uintptr(cgroup_id), uintptr(quotaMega), 0)
	new_quota := uint64(update_quota)

	if new_quota == 1 {
		log.Printf("Quota Set Failed\n")
		return new_quota, 1
	} 

	log.Printf("Quota Set Success!\n")
	log.Printf("ret: %d\n", new_quota)
	return uint64(new_quota), 0
}

func handle_mem_req(cgroup_id uint64) (uint64) {
	log.Printf("cgroup_id: %d\n", cgroup_id)
	avail_mem_ret, _, _ := syscall.Syscall(RESIZE_MEM_SYSCALL, uintptr(cgroup_id), 0, 0)
	avail_mem := uint64(avail_mem_ret)

	log.Printf("[INFO]: EC Agent: Reclaimed memory is: %d\n", avail_mem)
	return avail_mem
}

func handle_resize_max_mem(cgroup_id uint64, new_limit uint64, is_memsw int) (uint64) {
	log.Printf("setting new mem limit to: %d\n", new_limit)
	avail_mem_ret, _, _ := syscall.Syscall(RESIZE_MEM_SYSCALL, uintptr(cgroup_id), uintptr(new_limit), uintptr(is_memsw))
	avail_mem := uint64(avail_mem_ret)

	if avail_mem == 0 {
		log.Printf("[INFO]: EC Agent: resize_max_mem fails\n", avail_mem)
	}
	return avail_mem
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
			log.Println("ERROR in reading Header: ", err.Error())
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
		log.Println("--------------- BEGIN NEW REQUEST ---------------")
		switch rxMsg.GetReqType() {
		case 0:
			log.Println("CPU Request")
			updated_quota, ret = handle_cpu_req(uint64(rxMsg.GetCgroupId()), rxMsg.GetQuota())
		case 1:
			log.Println("Memory Request")
			ret = handle_mem_req(uint64(rxMsg.GetCgroupId()))
		case 2:
			log.Println("Init Request")
		case 3:
			log.Println("CPU SLICE")
		case 4: 
			container_id, ret = connect_container(rxMsg.GetClientIp(), rxMsg.GetPayloadString())
			if ret != 0 {
				log.Println("[ERROR] Initial Container Connection failed...")
				log.Println(container_id)
			}
		case 5:
			log.Println("Handle RESIZE MAX/MIN")
			ret = handle_resize_max_mem(uint64(rxMsg.GetCgroupId()), rxMsg.GetRsrcAmnt(), 0)
		default:
			log.Println("[ERROR] Not going in the right way! request type is invalid!")
		}
		log.Println("--------------- END NEW REQUEST ---------------")

		//log.Println("Docker Container id:", container_id)
		//log.Println("Updated Quota", updated_quota)
		txMsg := &msg_struct.ECMessage{
			ReqType: proto.Int32(rxMsg.GetReqType()),
			RsrcAmnt: proto.Uint64(ret),
			Quota: proto.Uint64(updated_quota),
			PayloadString: proto.String(container_id),
			Request: proto.Uint32(rxMsg.GetRequest()),
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

func main() {
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		if conn, err := l.Accept(); err == nil {
			go handleConnection(conn)
		}
	}
}