package main

import "fmt"

type Hello interface {
}

func main() {
	//if方便写法
/*	var a interface{}
	a = 111
	if _, ok := a.(*int); ok {
		fmt.Println("ok")
	}
*/
	var x Hello
	x = "hello"
	//类型断言
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case []int:
		fmt.Println("[]int")
	case []string:
		fmt.Println("[]string")
	default:
		fmt.Println("Don't know!")
	}
}
