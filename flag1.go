//flag包的使用，其实就是获取命令行参数，也可以用os.Args函数获取一个参数数组，但是没有 -flagname value的形式，类似php
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port = flag.Int("p", 8080, "Http server listen port")  //第一个参数-flagname参数名，第二个默认值，第三个提示信息
	var help = flag.Bool("h", false, "help")
	flag.Parse()  //必须写这个，不然无效
	if *help {  //如果有-p就显示提示信息
		flag.Usage()
	}

	http.HandleFunc("/", home)
	fmt.Printf("Http server is run at 0.0.0.0:%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
}
