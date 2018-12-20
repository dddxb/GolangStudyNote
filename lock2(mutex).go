/*Mutex（互斥锁）

Mutex 为互斥锁，Lock() 加锁，Unlock() 解锁
在一个 goroutine 获得 Mutex 后，其他 goroutine 只能等到这个 goroutine 释放该 Mutex
使用 Lock() 加锁后，不能再继续对其加锁，直到利用 Unlock() 解锁后才能再加锁
在 Lock() 之前使用 Unlock() 会导致 panic 异常
已经锁定的 Mutex 并不与特定的 goroutine 相关联，这样可以利用一个 goroutine 对其加锁，再利用其他 goroutine 对其解锁
在同一个 goroutine 中的 Mutex 解锁之前再次进行加锁，会导致死锁
适用于读写不确定，并且只有一个读或者写的场景
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var balance int
var mu sync.Mutex

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance += amount
}

func Withdraw(amount int) {
	mu.Lock()
	defer mu.Unlock()
	Deposit(-amount)
	if balance < 0 {
		Deposit(amount)
	}
}

func main() {
	for i := 0; i < 10000; i++ {
		go Deposit(i)
		go Withdraw(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(balance)
}
