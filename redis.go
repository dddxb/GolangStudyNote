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
	Err(err)
	val, err := r.Get("a").Result()
	Err(err)
	fmt.Println(val)
}

func Err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
