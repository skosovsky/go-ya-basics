package main

import (
	"log"
	"strings"
)

func Mul(a any, b int) any {
	switch data := a.(type) {
	case int:
		return data * b
	case string:
		return strings.Repeat(data, b)
	default:
		return nil
	}
}

func main() {
	log.Println(Mul(1, 2))    //nolint:mnd // example
	log.Println(Mul("ха", 2)) //nolint:mnd // example
}
