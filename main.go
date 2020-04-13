package main

import (
	"log"
	"net"
	"bufio"
	"io"
	"github.com/golang/protobuf/proto"
	"msg_struct"
)

const MAXGCMNO = 30
const PORT = ":4445"
const BUFFSIZE = 2048

func handleConnection(conn net.Conn) {
	log.Printf("[dbg] Server: New fd created for new connection. Serving %s\n", conn.RemoteAddr().String())
	for {
		buff := make([]byte, BUFFSIZE)
		c := bufio.NewReader(conn)

		// read a single byte which contains the message length at the beginning of the message
		size, err := c.ReadByte()
		if err != nil {
			log.Println("ERROR in reading Header: ", err.Error())
		}
		log.Println("ProtoBuf Message Body length: ", size)
		
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
		// Debug statement
		// log.Println("Recieved message req type: ", rxMsg.GetReqType())
		switch rxMsg.GetReqType() {
		case 0:
			log.Println("CPU Request")
		case 1:
			log.Println("Memory Request")
		case 2:
			log.Println("Init Request")
		case 3:
			log.Println("CPU SLICE")
		case 4: 
			log.Println("CONNECT CONTAINER")
		case 5:
			log.Println("Handle RESIZE MAX/MIN")
		default:
			log.Println("[ERROR] Not going in the right way! request type is invalid!")
		}

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