package main

import "fmt"

type Point struct {
	x, y int
}

type Circle struct {
	Point
	r float32
}

func (p *Point) set(x, y int) {
	p.x = x
	p.y = y
}

func main() {
	var z Point
	z.set(100,100)
	fmt.Println(z)
}
