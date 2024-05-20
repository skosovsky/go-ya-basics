package main

import (
	"encoding/json"
	"log"
	"time"
)

type Person struct {
	Name        string    `json:"имя"`
	Email       string    `json:"почта"`
	DateOfBirth time.Time `json:"-"`
}

func main() {
	alex := Person{
		Name:        "Alex",
		Email:       "alex@yandex.ru",
		DateOfBirth: time.Time{},
	}

	jsonData, err := json.Marshal(alex)
	if err != nil {
		log.Println(err)
	}

	log.Printf("Man %v", string(jsonData))
}
