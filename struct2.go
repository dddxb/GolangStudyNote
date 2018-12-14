package main

import "fmt"

type Peopel struct {
	a int
}

type Man struct {
	Peopel
	b int
}

type Child struct {
	Man
	c int
}

func main() {
	x := Child{}
	x.a = 1
	x.b = 2
	x.c = 3
	fmt.Println(x)
}