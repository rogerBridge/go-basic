package main

import (
	"fmt"
)

func whoAmI(i interface{}) {
	switch i.(type) {
	case bool:
		fmt.Println("I am a bool")
	case int, float64:
		fmt.Println("I am a number")
	case string:
		fmt.Println("I am a string")
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	//whoAmI("hello")
	//whoAmI(1234)
	//if n, ok := strconv.Atoi("H"); ok != nil {
	//	log.Println("转化string到int的过程中出了点问题 :)", ok)
	//} else {
	//	fmt.Println(n)
	//}
	//
	//if len("123") == 4 {
	//	return
	//}
	//fmt.Println("Hello World :)")
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")

	}
}
