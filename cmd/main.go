package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a []int

	getType(a)
}

func getType(a interface{}) {

	switch reflect.TypeOf(a.([]interface{})).Kind() {
	case reflect.Slice:
		fmt.Println("slice")
	case reflect.Array:
		fmt.Println("array")
	case reflect.String:
		fmt.Println("string")
	default:
		fmt.Println("unknown type")
	}
}
