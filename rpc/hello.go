//rpc服务
package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *[]int) error {
	*reply = []int{1, 2, 3, 4, 5}
	return nil
}

type GameService struct{}

func (p *GameService) Login(request string, reply *int) error {
	result := 0
	if request == "login" {
		result = 1
	}
	*reply = result
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	rpc.RegisterName("GameService", new(GameService))
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := listen.Accept()
	if err != nil {
		log.Fatal(err)
	}
	rpc.ServeConn(conn)
}
