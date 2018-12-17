//创建一个Test接口，并创建一个Example类型实现Test接口
package main

import (
	"fmt"
	"io"
)

type Example int

type Test interface {
	//Writer(p []byte) (int, error)
	//Reader(p []byte) (int, error)
	//Close() error
	//或者
	io.Writer
	io.Reader
	io.Closer
}

func (e *Example) Write(p []byte) (int, error) {
	fmt.Printf("%s\n",p)
	return 1, nil
}

func (e *Example) Read(p []byte) (int, error) {
	fmt.Printf("%s\n",p)
	return 1, nil
}

func (e *Example) Close() error {
	fmt.Printf("Example is closed!\n")
	return nil
}

func main() {
	var e Example
	testFunc(&e)
}

func testFunc(t Test) {
	t.Read([]byte("Example is writing!"))
	t.Write([]byte("Example is reading!"))
	t.Close()
}
