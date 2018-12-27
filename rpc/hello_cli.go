//调用rpc服务
package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply int
	err = client.Call("GameService.Login", 1, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}