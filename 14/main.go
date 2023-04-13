package main

import (
	"fmt"
	"reflect"
)

func main() {
	var some_var interface{}

	some_var = "hello"

	switch some_var.(type) {
	case int:
		fmt.Println("Type is int")
	case string:
		fmt.Println("Type is string")
	case bool:
		fmt.Println("Type is bool")
	default:
		if reflect.TypeOf(some_var).Kind() == reflect.Chan {
			fmt.Println("Type is channel")
		}
	}
}
