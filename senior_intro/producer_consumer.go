//并发：生产者与消费者模型
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Producer(x int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * x
	}
}

func Consumer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	ch := make(chan int, 64)
	go Producer(1, ch)
	go Producer(2, ch)
	go Producer(3, ch)

	go Consumer(ch)

	// Ctrl+C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
