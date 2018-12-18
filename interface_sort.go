package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	TypeGet()
	StrSort()
	IntSort()
	time.Sleep(1*time.Second)
}

// 数字排序
type IntSlice []int
func (i IntSlice) Len() int           { return len(i) }
func (i IntSlice) Less(x, y int) bool { return i[x] < i[y] }
func (i IntSlice) Swap(x, y int)      { i[x], i[y] = i[y], i[x] }
func IntSort() {
	x := IntSlice{1, 5, 2, 3, 11, 4, 33, 6, 21, 22, 11, 4, 2, 15, 8, 3, 6, 17, 12, 55, 43, 22, 13}
	sort.Sort(x)
	fmt.Println(x)
}

// 排序字符串
type StringSlice []string
func (s StringSlice) Len() int           { return len(s) }
func (s StringSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s StringSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func StrSort() {
	x := StringSlice{"ah", "aa", "as", "cs", "bc", "ba"}
	sort.Sort(x)
	fmt.Println(x)
}

// %T获取数据类型
func TypeGet() {
	var a interface{} = 1
	types := fmt.Sprintf("%T", a)
	fmt.Printf("%s\n",types)
}
