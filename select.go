//select多路复用
package main

import "fmt"

func main() {
	a := make(chan byte)
	b := make(chan byte)
	go test(a, 'a')
	go test(b, 'b')
	for {
		select {
		case x := <-a:
			fmt.Printf("%s\n", string(x))
		case y := <-b:
			fmt.Printf("%s\n", string(y))
		}
	}
}

func test(c chan byte, b byte) {
	for i := 1; i < 1000; i++ {
		c <- b
	}
	close(c)
}
