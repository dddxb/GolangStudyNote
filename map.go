package main

import (
	"fmt"
	"sort"
)

func main() {
	test3()
}

//创建map，赋值，取值
func test1() {
	arg1 := make(map[string]int)
	arg2 := map[string]int{"china": 36, "japan": 9}
	arg1["canada"] = 18
	delete(arg2, "japan")
	fmt.Println(arg1, arg2, arg1["none"])

	for k, v := range arg2 {
		fmt.Println(k, v)
	}
}

//循环map顺序不是固定的，利用sort排序
func test2() {
	args := map[string]int{
		"a": 1,
		"c": 2,
		"b": 1,
		"d": 4,
		"e": 9,
		"f": 1,
		"g": 0,
		"h": 4}
	var names []string
	for name, _ := range args {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Println(args[name])
	}
}

//map数组
func test3() {
	type Country map[string]int
	m := Country{"china": 99}
	s := []Country{m, m, m}
	fmt.Println(s)
}
