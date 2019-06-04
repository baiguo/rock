package network

import (
	"fmt"
	"github.com/baiguo/rock/console"
	"log"
	"net"
	"strconv"
)

type TCPServer struct {
	ln              net.Listener
}

func (server *TCPServer) Init()  {
	listener, err := net.Listen("tcp", ":"+ strconv.Itoa(*console.Port))
	if err != nil{
		log.Fatal("this is TCPServer listen fatal err: ",err)
		return
	}

	log.Printf("this is TCPServer listen to prot: %d", *console.Port)
	server.ln = listener

}

func (server *TCPServer) Process(conn net.Conn){
	defer conn.Close()

	for{
		bytes := make([]byte,1024)
		n, err := conn.Read(bytes)
		if err == nil{
			fmt.Println(string(bytes[0:n]))
		}else{
			fmt.Println("read err:",err)
			continue
		}
	}
}

func (server *TCPServer) Run()  {

	for{
		conn, err := server.ln.Accept()
		if err != nil {
			log.Print("server.ln.Accept err: %v", err)
			continue
		}

		go server.Process(conn)
	}

}
