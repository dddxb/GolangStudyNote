package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	password := "panco123456"
	w := md5.New()
	io.WriteString(w, password)
	token := fmt.Sprintf("%x", w.Sum(nil))
	fmt.Println(token)
}
