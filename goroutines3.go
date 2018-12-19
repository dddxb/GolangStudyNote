//并发Clock客户端
package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	//模拟多个客户端连接服务端
	for i := 0; i < 10; i++ {
		go connectServer("127.0.0.1:9999")
	}
	time.Sleep(3600 * time.Minute)
}

func connectServer(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn) //利用io.Copy将来buffer传输到屏幕
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
