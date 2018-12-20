//goroutines mutex互斥锁，就是一个协程Lock住，另一个协程必须等Unlock才可以继续
//这个程序Desposit和Balance会争抢锁，所以结果不一定是什么时候打印出balance
package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var balance int

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance += amount
}

func Balance()int {
	mu.Lock()
	defer mu.Unlock()
	Deposit(-amount)
	if balance < 0 {
		Deposit(amount)
	}
}

func view() {
	for {
		for _,v := range `-\|/` {
			fmt.Printf("\r%c", v)
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	go view()
	for i := 0; i < 5; i++ {
		go Deposit(10)
	}
	go Balance()
	time.Sleep(10 * time.Second)
}
