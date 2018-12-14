package main

import "fmt"

func main() {
	slice := make([]int, 1)
	reserve(slice)
	slice = append(slice, 20)
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))

	array := [3]int{1, 2, 3}
	updateArray(array)
	fmt.Println(array)
}

//silence包含隐式指针，不传指针赋值会改变原有silence
func reserve(s []int) {
	s[0] = 2121
}

//array不会，array必须传指针才会改变原有array
func updateArray(arr [3]int) {
	arr[1] = 999
}
