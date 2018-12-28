//atomic原子操作：atomic.AddInt64函数调用保证了total的读取、更新和保存是一个原子操作，因此在多线程中访问也是安全的。
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total int64

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&total, 1)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(total)
}
