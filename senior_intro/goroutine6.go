/*首先，我们创建了一个带缓存的管道，管道的缓存数目要足够大，保证不会因为缓存的容量引起不必要的阻塞。然后我们开启了多个后台线程，分别向不同的搜索引擎提交搜索请求。当任意一个搜索引擎最先有结果之后，都会马上将结果发到管道中（因为管道带了足够的缓存，这个过程不会阻塞）。但是最终我们只从管道取第一个结果，也就是最先返回的结果。

通过适当开启一些冗余的线程，尝试用不同途径去解决同样的问题，最终以赢者为王的方式提升了程序的相应性能。
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 32)

	go func() {
		ch <- searchByBing("golang")
	}()
	go func() {
		ch <- searchByGoogle("golang")
	}()
	go func() {
		ch <- searchByBaidu("golang")
	}()

	fmt.Println(<-ch)
}

func searchByBaidu(a string) string {
	time.Sleep(1 * time.Millisecond)
	return "a"
}
func searchByBing(a string) string {
	time.Sleep(1 * time.Millisecond)
	return "b"
}
func searchByGoogle(a string) string {
	time.Sleep(1 * time.Millisecond)
	return "c"
}
