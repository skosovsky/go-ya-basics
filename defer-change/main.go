package main

import "fmt"

var Global = 5 //nolint:gochecknoglobals // example

func main() {
	fmt.Println(Global) //nolint:forbidigo // example
	withDefer()
	fmt.Println(Global) //nolint:forbidigo // example
	UseGlobal()
	fmt.Println(Global) //nolint:forbidigo // example
}

// withDefer Просто изменяем переменную внутри deferю.
func withDefer() {
	defer func() {
		reserveGlobal := Global
		Global = 10
		fmt.Println(Global) //nolint:forbidigo // example
		Global = reserveGlobal
	}()
}

// UseGlobal Делаем бэкап переменной и возвращаемся к нему в defer.
func UseGlobal() {
	defer func(checkout int) {
		Global = checkout // присваиваем Global значение аргумента
	}(Global) // копируем значение Global в аргументы функции
	Global = 42         // Изменяем Global
	fmt.Println(Global) //nolint:forbidigo // example
	// Здесь будет вызвана отложенная функция
}
