//select多路复用
package main

import (
	"fmt"
)

func main() {
	a := make(chan byte)
	b := make(chan byte)
	done := make(chan struct{})
	go test(a, 'a', done)
	go test(b, 'b', done)
	for {
		select {
		case x := <-a:
			fmt.Println(string(x))
		case y := <-b:
			fmt.Println(string(y))
		case <-done:
			return
		default:
		}
	}
}

func test(c chan byte, b byte, done chan struct{}) {
	for i := 1; i < 1000; i++ {
		c <- b
	}
	close(c)
	done <- struct{}{}
}
