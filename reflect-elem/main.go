package main

import (
	"log"
	"reflect"
)

func base() {
	var varBool *bool
	log.Println(reflect.ValueOf(varBool).IsNil()) // true

	trueVal := true
	reflect.ValueOf(&varBool).Elem().Set(reflect.ValueOf(&trueVal))

	log.Println(reflect.ValueOf(varBool).IsNil())       // false
	log.Println(reflect.ValueOf(varBool).Elem().Bool()) // true — получить значение через рефлексию
	log.Println(*varBool)                               // true — получить значение без рефлексии

	reflect.ValueOf(varBool).Elem().Set(reflect.ValueOf(false))
	log.Println(*varBool)

	reflect.ValueOf(varBool).Elem().SetBool(true)
	log.Println(*varBool)

	varInt := 100
	varIntValue := reflect.ValueOf(varInt)
	log.Println(varIntValue.IsZero()) // false
	log.Println(varIntValue.Int())    // 100

	var varPtrInt *int
	varPtrIntValue := reflect.ValueOf(varPtrInt)
	log.Println(varPtrIntValue.IsNil())  // true
	log.Println(varPtrIntValue.IsZero()) // true
}

func typeKind() {
	var varBool *bool
	log.Println(reflect.ValueOf(varBool).Kind()) // ptr
	log.Println(reflect.ValueOf(varBool).Type()) // *bool

	var varFloat float32
	log.Println(reflect.ValueOf(varFloat).Kind()) // float32
	log.Println(reflect.ValueOf(varFloat).Type()) // float32

	var varMap map[string]int
	log.Println(reflect.ValueOf(varMap).Kind()) // map
	log.Println(reflect.ValueOf(varMap).Type()) // map[string]int

	varStruct := struct{ Value int }{}             //nolint:exhaustruct  // example
	log.Println(reflect.ValueOf(varStruct).Kind()) // struct
	log.Println(reflect.ValueOf(varStruct).Type()) // struct { Value int }
}

func isEqual(x, y any) bool {
	return reflect.DeepEqual(x, y)
}

func isNil(x any) bool {
	if x == nil {
		return true
	}

	value := reflect.ValueOf(x)
	if value.Kind() == reflect.Ptr {
		return false
	}

	if value.IsNil() {
		return true
	}

	return false
}

func main() {
	base()

	log.Println()

	typeKind()

	log.Println(isEqual(reflect.Bool, reflect.Bool)) // true

	log.Println()

	log.Println(isNil(nil)) // true
}
