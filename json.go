package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Actor struct {
	Name string
	Age  int
}

type Movie struct {
	Title  string
	Year   int
	Color  bool
	Actors []Actor
}

func main() {
	movies := []Movie{
		{
			Title:  "功夫",
			Year:   1998,
			Color:  true,
			Actors: []Actor{Actor{Name:"成龙", Age: 53}, Actor{Name: "王宝强", Age: 33}},
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
}
