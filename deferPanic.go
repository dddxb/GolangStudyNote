package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	defer fmt.Printf("a")
	defer fmt.Printf("b")

	fmt.Printf("c")
	fmt.Printf("d")

	test9()
	return
}

func test9() {
	defer printStack()
	panic("\nThis is a panic!")
}

//轷出堆栈俆息
func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
