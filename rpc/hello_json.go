//rpc服务
package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *[]int) error {
	*reply = []int{1, 2, 3, 4, 5}
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
