package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 42
	v := reflect.ValueOf(x)

	fmt.Println("Type:", v.Type())
	fmt.Println("Kind:", v.Kind())
	fmt.Println("Value:", v.Int()) // Retrieve the integer value
}
