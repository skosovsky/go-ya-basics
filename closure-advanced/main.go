package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// countCall — функция-обёртка для подсчёта вызовов
func countCall(f func(string)) func(string) {
	cnt := 0
	// Получаем имя функции. Подробнее об этом вызове будет рассказано в следующем курсе.
	funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return func(s string) {
		cnt++
		fmt.Printf("Функция %s вызвана %d раз\n", funcName, cnt)
		f(s)
	}
}

// metricTimeCall — функция-обёртка для замера времени
func metricTimeCall(f func(string)) func(string) {
	return func(s string) {
		start := time.Now() // получаем текущее время
		f(s)
		fmt.Println("Execution time: ", time.Since(start)) // получаем интервал времени как разницу между двумя временными метками
	}
}

func myPrint(s string) {
	fmt.Println(s)
}

func main() {
	countedPrint := countCall(myPrint)
	countedPrint("Hello world")
	countedPrint("Hi")

	// Обратите внимание, что мы оборачиваем уже обёрнутую функцию, а значение счётчика вызовов при этом сохраняется.
	countAndMetricPrint := metricTimeCall(countedPrint)
	countAndMetricPrint("Hello")
	countAndMetricPrint("World")

}
