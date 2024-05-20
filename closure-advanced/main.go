package main

import (
	"log"
	"reflect"
	"runtime"
	"time"
)

// countCall — функция-обёртка для подсчёта вызовов.
func countCall(fn func(string)) func(string) {
	cnt := 0
	// Получаем имя функции. Подробнее об этом вызове будет рассказано в следующем курсе.
	funcName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()

	return func(s string) {
		cnt++
		log.Printf("Функция %s вызвана %d раз\n", funcName, cnt)
		fn(s)
	}
}

// metricTimeCall — функция-обёртка для замера времени.
func metricTimeCall(fn func(string)) func(string) {
	return func(s string) {
		start := time.Now() // получаем текущее время
		fn(s)
		log.Println("Execution time: ", time.Since(start)) // получаем интервал времени как разницу между двумя временными метками
	}
}

func myPrint(s string) {
	log.Println(s)
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
