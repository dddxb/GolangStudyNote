/**
复合 数据类型函数传值测试
测试结果：
	1、array、struct传本身不是指针
	2、slice、map穿本身也会变成指针
 */
package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := []int{1, 2, 3}
	c := map[string]int{"a": 1, "b": 2}
	d := struct {
		Name int
	}{Name: 1}

	t1(&a)
	t2(b)
	t3(c)
	t4(&d)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

func t1(a *[3]int) {
	a[0] = 9
}

func t2(b []int) {
	b[0] = 9
}

func t3(c map[string]int) {
	c["a"] = 9
}

func t4(d *struct{Name int}) {
	d.Name = 9
}
