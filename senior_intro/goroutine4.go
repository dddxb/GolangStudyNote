package main

import (
	"fmt"
	"sync"
)

func main() {
	main2()
	var done = make(chan bool)
	var a string
	go func() {
		a = "Hello World!"
		done <- true
	}()

	<-done
	fmt.Println(a)
}

//可以确定后台线程的mu.Unlock()必然在println("你好, 世界")完成后发生（同一个线程满足顺序一致性），main函数的第二个mu.Lock()必然在后台线程的mu.Unlock()之后发生（sync.Mutex保证），此时后台线程的打印工作已经顺利完成了。
func main2() {
	var mu sync.Mutex
	mu.Lock()
	var a string

	go func() {
		a = "Hello World!"
		mu.Unlock()
	}()
	mu.Lock()

	fmt.Println(a)
}
