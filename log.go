package main

import (
	"log"
	"os"
	"time"
)

func main() {
	go WriteLog("hello")
	time.Sleep(1 * time.Second)
}

//记录日志，可追加
func WriteLog(s string) {
	file, err := os.OpenFile("./log.log", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		log.Fatal("Not Access!")
	}
	offset, err := file.Seek(0, os.SEEK_END)
	file.WriteAt([]byte(time.Now().Format("2006/1/2 15:04:05")+"："+s+"\n"), offset)
	file.Close()
}
