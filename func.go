package main

import (
	"fmt"
)

func main() {
	x := unknowFunc()
	fmt.Println(x())
	fmt.Println(x())

	fmt.Println(changeFunc(1, "a", 2))
}

//匿名函数会保存函数中变量状态
func unknowFunc() func() int {
	x := 1
	return func() int {
		x++
		return x * x
	}
}

//可变参数函数
func changeFunc(x ...interface{}) []interface{} {
	var z []interface{}
	for _, v := range x {
		z = append(z, v)
	}
	return z
}
