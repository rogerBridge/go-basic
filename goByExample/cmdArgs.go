package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] // without program
	fmt.Println(args) // if I need 2 arguments, then len(args) must be 2
	if len(args) == 2 {
		fmt.Println(args[0], args[1])
	}else {
		panic("len(args) is not 2")
	}
}
