package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

func main() {
	alex := Person{
		Name:  "Alex",
		Email: "alex@yandex.ru",
	}

	jsonData, err := json.Marshal(alex)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Man %v", string(jsonData)) //nolint:forbidigo // it's learning code
}
