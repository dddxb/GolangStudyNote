//安全的并发修改数据
package main

import (
	"fmt"
	"time"
)

var deposits = make(chan int) //传输余额变化
var balances = make(chan int) //存储当前余额

//修改余额
func Deposit(amount int) {
	deposits <- amount
}

//当前余额
func Balance() int {
	return <-balances
}

//修改余额监听
func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

//主函数
func main() {
	go teller()
	for i := 1; i < 10000; i++ {
		go Deposit(1)
		go Deposit(-1)
	}
	time.Sleep(5 * time.Second)
	fmt.Println(Balance())
}
