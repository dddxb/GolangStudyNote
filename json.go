package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Actor struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Movie struct {
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Color  bool    `json:"color"`
	Actors []Actor `json:"actors"`
}

func main() {
	//结构体转json
	movies := []Movie{
		{
			Title:  "功夫",
			Year:   1998,
			Color:  true,
			Actors: []Actor{Actor{Name: "成龙", Age: 53}, Actor{Name: "王宝强", Age: 33}},
		},
		{
			Title:  "僵尸先生",
			Year:   1999,
			Color:  true,
			Actors: []Actor{Actor{Name: "林正英", Age: 43}, Actor{Name: "范冰冰", Age: 33}},
		},
	}
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)

	//json转结构体
	var data2 []Movie
	err2 := json.Unmarshal(data, &data2)
	if err != nil {
		log.Fatal(err2)
	}
	fmt.Println(data2)
}
