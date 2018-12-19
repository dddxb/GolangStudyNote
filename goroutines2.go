//并发Clock服务器
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const PORT = "9999"

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server is run at port " + PORT + "\n")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	fmt.Printf("%s is connect!\n", c.RemoteAddr())
	defer fmt.Printf("%s is disconnect!\n", c.RemoteAddr())
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("01-02 03:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
