package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go send(ch)
	get(ch)
}

func send(ch chan int) {
	ch <- 1 // 将999发送给channel
}

func get(ch chan int) {
	x := <- ch // 接收channel并赋值给x
	fmt.Println(x)
}
