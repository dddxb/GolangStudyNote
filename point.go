package main

import (
	"fmt"
	"time"
)

type Peopel struct {
	name string
	age  int
}

func main() {
	x := 0
	for z := 0; z < 10000; z++ {
		go incr(&x, 1)
	}
	time.Sleep(2*time.Second)
	fmt.Println(x)

}

func incr(x *int, num int) {
	*x = *x + num
}
