//带缓存的channel
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	response := make(chan string, 3)
	go func() { response <- request("http://www.baidu.com/1.html") }()
	go func() { response <- request("http://www.baidu.com/2.html") }()
	go func() { response <- request("http://www.baidu.com/3.html") }()
	x := <-response
	fmt.Println(x)
}

func request(address string) (response string) {
	num := Generate_Randnum()
	time.Sleep(time.Duration(num) * time.Second)
	return address
}

//随机数
func Generate_Randnum() float32{
	rand.Seed(time.Now().Unix())
	rnd := rand.Float32()
	return rnd
}