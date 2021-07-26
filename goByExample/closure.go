package main

import (
	"fmt"
)

// 一个函数A, 含有一个变量i, 返回一个函数B, B可以操作i, 此时, A+B就形成了闭包
// 使用方法, 将A赋值给一个变量名a, 一个变量名b, a()的所有操作和b()的所有操作是互不影响的
func increseInt() func() int {
	i := 0
	return func() int {
		i += 1
		//pc, _, _, _ := runtime.Caller(1)
		//fmt.Println(runtime.FuncForPC(pc).Name())
		return i
	}
}

func main() {
	nextInt := increseInt()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	nextInt1 := increseInt()
	fmt.Println(nextInt1())
}

