package main

import (
	"log"
)

func Generate(seed int) func() {
	return func() {
		log.Println(seed) // замыкание получает внешнюю переменную seed
		seed += 2         // переменная модифицируется
	}
}

func fib() func() int {
	x1, x2 := 0, 1
	// Возвращаемая функция замыкает x1, x2.
	return func() int {
		x1, x2 = x2, x1+x2

		return x1
	}
}

func main() {
	log.Println("iterator")
	iterator := Generate(0)
	iterator()
	iterator()
	iterator()
	log.Println()

	log.Println("fibonacci")
	fibonacci := fib()       // Получили функцию-замыкание. fibonacci() — захватила x1, x2, x1 = 0, x2 = 1.
	log.Println(fibonacci()) // x1 = 1, x2 = 1
	log.Println(fibonacci()) // x1 = 1, x2 = 2
	log.Println(fibonacci()) // x1 = 2, x2 = 3
	log.Println(fibonacci()) // x1 = 3, x2 = 5
	log.Println(fibonacci()) // x1 = 5, x2 = 8
	log.Println(fibonacci()) // x1 = 8, x2 = 13
}
