package main

import "fmt"

func Generate(seed int) func() {
	return func() {
		fmt.Println(seed) //nolint:forbidigo // замыкание получает внешнюю переменную seed
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
	fmt.Println("iterator")
	iterator := Generate(0)
	iterator()
	iterator()
	iterator()
	fmt.Println()

	fmt.Println("fibonacci")
	fibonacci := fib()       // Получили функцию-замыкание. fibonacci() — захватила x1, x2, x1 = 0, x2 = 1.
	fmt.Println(fibonacci()) //nolint:forbidigo // x1 = 1, x2 = 1
	fmt.Println(fibonacci()) //nolint:forbidigo // x1 = 1, x2 = 2
	fmt.Println(fibonacci()) //nolint:forbidigo // x1 = 2, x2 = 3
	fmt.Println(fibonacci()) //nolint:forbidigo // x1 = 3, x2 = 5
	fmt.Println(fibonacci()) //nolint:forbidigo // x1 = 5, x2 = 8
	fmt.Println(fibonacci()) //nolint:forbidigo // x1 = 8, x2 = 13
}
