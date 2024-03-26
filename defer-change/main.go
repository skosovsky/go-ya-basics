package main

import "fmt"

var Global = 5 //nolint:gochecknoglobals // it's learning code

func main() {
	fmt.Println(Global) //nolint:forbidigo // it's learning code
	withDefer()
	fmt.Println(Global) //nolint:forbidigo // it's learning code
	UseGlobal()
	fmt.Println(Global) //nolint:forbidigo // it's learning code
}

// withDefer Просто изменяем переменную внутри deferю
func withDefer() {
	defer func() {
		reserveGlobal := Global
		Global = 10
		fmt.Println(Global) //nolint:forbidigo // it's learning code
		Global = reserveGlobal
	}()
}

// UseGlobal Делаем бэкап переменной и возвращаемся к нему в defer.
func UseGlobal() {
	defer func(checkout int) {
		Global = checkout // присваиваем Global значение аргумента
	}(Global) // копируем значение Global в аргументы функции
	Global = 42         // Изменяем Global
	fmt.Println(Global) //nolint:forbidigo // it's learning code
	// Здесь будет вызвана отложенная функция
}
