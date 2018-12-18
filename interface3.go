//利用空interface实现任意类型
package main

import "fmt"

type Any interface{}

func main() {
	var x Any
	x = 1
	x = true
	x = "hello"
	x = []byte("heihei")
	x = []int{1, 2, 3, 4}
	x = struct {
		s int
	}{}
	x = map[string]string{"map": "hello"}
	x = 1
	fmt.Printf("%d", x)
}
