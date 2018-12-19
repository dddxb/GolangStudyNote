// 单方向channel参数定义，只接受的不能close
// 类垄 chan<- int 表示一
//丧叧収送 int 癿 channel，叧能収送丌能掍收。相反，类垄<-chan int 表示一
//丧叧掍收 int 癿 channel，叧能掍收丌能収送。
package main

import "fmt"

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
