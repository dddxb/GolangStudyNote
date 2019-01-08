package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var r *redis.Client

func connect() {
	r = redis.NewClient(&redis.Options{
		Addr:     "192.168.5.128:6379",
		Password: "0825",
		DB:       0,
	})
	_, err := r.Ping().Result()
	Err(err)
}

func main() {
	connect()
	err := r.Set("a", "a", 0).Err()
	r.HMSet("player1", map[string]interface{}{
		"name": "panco",
		"age":  23,
	})
	Err(err)
	cmd, err2 := r.HMGet("player1", "name", "age").Result()
	Err(err2)
	fmt.Println(cmd[0])
}

func Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
