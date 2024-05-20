package main

import (
	"reflect"
	"strconv"
)

func ChangeFieldByName(val any, fieldName string, newVal int) {
	value := reflect.ValueOf(val)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return
	}

	field := value.FieldByName(fieldName)
	if field.IsValid() {
		if field.CanSet() {
			switch field.Kind() { //nolint:exhaustive // example
			case reflect.Int:
				field.SetInt(int64(newVal))
			case reflect.String:
				field.SetString(strconv.Itoa(newVal))
			default:
				panic("unhandled default case")
			}
		}
	}
}

func main() {

}
