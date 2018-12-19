//三个协程之间通信
package main

import (
	"fmt"
)

func main() {
	counter := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch1 := make(chan []int)
	ch2 := make(chan int)
	done := make(chan struct{})
	go cal(ch1, ch2)
	ch1 <- counter  //同一个协程之间不能通信
	go prints(ch2, done)
	<-done
}

func cal(ch1 chan []int, ch2 chan int) {
	a := <-ch1
	x := 0
	for _, v := range a {
		x += v
	}
	x *= x
	ch2 <- x
}

func prints(ch2 chan int, done chan struct{}) {
	number := <-ch2
	fmt.Println(number)
	done <- struct{}{}
}
