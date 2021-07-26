package main

import (
	"fmt"
	"reflect"
)

func main() {
	//a := "安瑞峰"
	//b := 24
	type Student struct {
		Age     int
		Sexual  int
		Name    string
		Chinese float64
		Math    float64
		English float64
	}
	b := &Student{24, 1, "leo", 80, 95, 86}
	fmt.Println(reflect.TypeOf(b), reflect.ValueOf(b), reflect.TypeOf(b))
}
