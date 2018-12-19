package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	done := make(chan struct{}) //用来同步，不让main退出

	go func() {
		for x := 0; ; x++ {
			if x > 1000 {
				break
			}
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	go func() {
		for {
			get, ok := <-squares
			if !ok {
				break
			}
			fmt.Println(get)
		}
		done <- struct{}{}  //结束main
	}()

	<-done //等待接受done
}
