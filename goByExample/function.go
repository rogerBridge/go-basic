package main

import (
	"reflect"
)

func Plus(a int, b int) (int, error) {
	if reflect.TypeOf(a) == reflect.TypeOf(b) && reflect.TypeOf(a).String() == "int" {
		return a+b, nil
	}
	var e error
	return 0, e
}

func main() {
	Plus(1,2)
}