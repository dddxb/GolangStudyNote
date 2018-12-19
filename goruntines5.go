//Echo客户端
package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	connectServer("127.0.0.1:9999")
	time.Sleep(3600 * time.Minute)
}

func connectServer(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)  //开一个协程接收服务器消息并打印，把conn的消息发送到屏幕输出
	mustCopy(conn, os.Stdin)  //把输入的东西发送到连接(服务器)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
