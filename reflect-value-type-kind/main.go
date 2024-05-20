package main

import (
	"log"
	"reflect"
)

func main() {
	var varBool *bool
	log.Println(reflect.ValueOf(varBool).Kind()) // ptr — указатель
	log.Println(reflect.ValueOf(varBool).Type()) // *bool — указатель на bool

	var varFloat float32
	log.Println(reflect.ValueOf(varFloat).Kind()) // float32
	log.Println(reflect.ValueOf(varFloat).Type()) // float32

	var varMap map[string]int
	log.Println(reflect.ValueOf(varMap).Kind()) // map
	log.Println(reflect.ValueOf(varMap).Type()) // map[string]int

	varStruct := struct{ Value int }{Value: 0}
	log.Println(reflect.ValueOf(varStruct).Kind()) // struct
	log.Println(reflect.ValueOf(varStruct).Type()) // struct { Value int }
}
