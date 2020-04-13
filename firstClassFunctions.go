package main

import "fmt"

// Pass functions as arguments to other functions
func playFunc(a func(a int, b int) int) {
	a(12, 13)
}

func playFuncAsArguments() {
	a := func(a int, b int) int{
		fmt.Println(a+b)
		return a+b
	}
	playFunc(a)
}
