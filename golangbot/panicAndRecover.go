package main

import (
	"fmt"
	"runtime/debug"
)

func playPanic() {
	firstName := "An"
	//fullName(&firstName, nil)
	fullName1(&firstName, nil)
}

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
}

func fullName1(firstName *string, lastName *string) {
	defer fmt.Println("call fullName1")
	if firstName == nil {
		panic("runtime error: firstName error")
	}
	if lastName == nil {
		panic("runtime error: lastname error")
	}
	fmt.Println("%s %s", firstName, lastName)
}

// 从panic中恢复过来
func playRecover() {
	defer r()
	panic("panic test")
}

// recover之后的处理过程
func r() {
	if r:=recover(); r!=nil {
		fmt.Println("Recovered ", r)
	}
	debug.PrintStack() // panic
}

