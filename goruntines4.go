//并发Echo服务器
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
	input := bufio.NewScanner(c)  //读取客户端消息
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)  //这个go可以使客户端并发处理同事发的消息，而不用等待处理完上一个消息
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
