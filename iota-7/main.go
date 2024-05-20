package main

import (
	"log"
)

const (
	one = 2*iota + 1
	three
	five
	seven
	nine
	eleven
)

func main() {
	log.Println(one, three, five, seven, nine, eleven)
}
