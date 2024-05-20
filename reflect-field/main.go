package main

import (
	"log"
	"reflect"
)

func ExtendedPrint(val any) {
	value := reflect.ValueOf(val)

	switch value.Kind() { //nolint:exhaustive // example
	case reflect.Ptr:
		if value.Elem().Kind() != reflect.Struct {
			log.Printf("Pinter to %v : %v", value.Elem().Type(), value.Elem())

			return
		}

		value = value.Elem()

	case reflect.Struct:
		log.Printf("All ok, %v", value)

	default:
		log.Printf("%v : %v", value.Type(), value)

		return
	}

	log.Printf("Struct of type %v and number of fields %d:\n", value.Type(), value.NumField())
	for idx := range value.NumField() {
		field := value.Field(idx)
		log.Printf("\tField %v : %v - val :%v\n", value.Type().Field(idx).Name, field.Type(), field)
	}
}

type myStruct struct {
	A int
	B string
	C bool
}

func main() {
	str := myStruct{
		A: 3, //nolint:mnd // example
		B: "some",
		C: false,
	}
	str2 := myStruct{
		A: 7, //nolint:mnd // example
		B: "some text",
		C: true,
	}

	ExtendedPrint(str)

	ExtendedPrint(str2)

	ExtendedPrint(struct {
		E int
		C string
	}{2, "some text"})

	ExtendedPrint("some string")
}
