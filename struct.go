package main

import "fmt"

type User struct {
	id       int64
	username string
	password string
	avatar   string
	gender   int8
	status   int8
}

func main() {
	test6()
}

func test4() {
	b := User{}
	b.id = 1
	b.username = "panco"
	b.password = "123456"
	b.avatar = "http://www.baidu.com/test.png"
	b.gender = 1
	b.status = 1

	a := User{2, "jack", "123456", "http://www.baidu.com/test.png", 1, 1}

	var c User
	c.id = 3
	c.username = "tom"
	c.password = "123456"
	c.avatar = "http://www.baidu.com/test.png"
	c.gender = 1
	c.status = 1

	fmt.Println(a, b, c)
}

func test5() {
	item := User{2, "jack", "123456", "http://www.baidu.com/test.png", 1, 1}
	arr := []User{item, item, item}
	fmt.Println(arr)
}

func test6() {
	a := User{2, "jack", "123456", "http://www.baidu.com/test.png", 1, 1}
	test7(&a)
	fmt.Println(a)

	b := User{2, "jack", "123456", "http://www.baidu.com/test.png", 1, 1}
	point := &b.id
	*point = 3
	fmt.Println(b)
}

func test7(a *User) {
	(*a).id = 99
}
