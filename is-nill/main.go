package main

import (
	"log"
	"reflect"
)

type MyType struct{}

func isNil(obj any) bool {
	if obj == nil {
		return true
	}

	objValue := reflect.ValueOf(obj)
	if objValue.Kind() == reflect.Ptr {
		return false
	}
	if objValue.IsNil() {
		return true
	}

	return false
}

func main() {
	var t *MyType
	log.Printf("Проверка типа (%v) на nil: %v\n", reflect.TypeOf(t), isNil(t))
}
