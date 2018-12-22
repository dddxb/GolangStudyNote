package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

var Online = make(map[*websocket.Conn]interface{}) //存储所有链接
var mu sync.RWMutex                                //读写锁

func main() {
	go printOnline()
	http.HandleFunc("/", wsHandle)
	http.ListenAndServe("0.0.0.0:9000", nil)
}

var upgrader = websocket.Upgrader{
	//允许ws跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//websocket
func wsHandle(w http.ResponseWriter, r *http.Request) {
	//检测用户链接是否是websocket
	if !websocket.IsWebSocketUpgrade(r) {
		fmt.Fprintf(w, "Please use websocket to connect this server!")
	} else {
		//如果是将http协议升级为websocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
		}
		go addOnline(conn)
		defer delOnline(conn)
		defer conn.Close()
		//循环接受消息
		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				break
				log.Println(err)
			}
			push(mt, message, conn)
		}
	}
}

//添加连接
func addOnline(conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	Online[conn] = map[string]interface{}{
		"conn": conn,
	}
}

//删除连接
func delOnline(conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	delete(Online, conn)
}

//打印连接map
func printOnline() {
	for {
		time.Sleep(5 * time.Second)
		fmt.Println(Online)
	}
}

//推送消息给所有人
func push(messageType int, message []byte, expect *websocket.Conn) {
	mu.RLock()
	defer mu.RUnlock()
	for k, _ := range Online {
		if k != expect {
			k.WriteMessage(messageType, message)
		}
	}
}
