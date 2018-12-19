//运算的同时显示加载图标
package main

import (
	"fmt"
	"time"
)

func main() {
	go loading()
	z := 1
	for i := 0; i < 1000000000; i++ {
		z += i*i*i*i /2/2/2/2/2/2*11111/222123
	}
	fmt.Printf("\r%d",z)
}

func loading() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(50 * time.Millisecond)
		}
	}
}
