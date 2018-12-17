//实现一个io.Writer接口类型，并且成功调用fmt.Fprintf方法
package main

import (
	"fmt"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0
	name := "Panco"
	x, _ := fmt.Fprintf(&c, "hello %s!", name)
	fmt.Println(x)
}