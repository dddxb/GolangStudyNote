package main

import (
	"fmt"
)

type Peopel int

func (p Peopel) hello(x int) {
	fmt.Println(x)
}

func main() {
	var a Peopel
	a.hello(1)
}
