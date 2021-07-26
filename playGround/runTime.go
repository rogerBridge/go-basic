package main

import (
	"encoding/hex"
	"fmt"
	"math"
)
func main(){
	//t0 := time.Now()
	//fmt.Println(runtime.GOOS)
	////fmt.Println(t0)
	//result := getSqrtSum(1.0, 10000000.0)
	////fmt.Println("结果是: ", result)
	//elapsed := time.Since(t0)
	//fmt.Println("😥", len("😥"))
	//fmt.Println("安", len("安"))
	//fmt.Println("a", len("a"))
	//fmt.Println(utf8.RuneCountInString("😥安瑞峰"))
	//fmt.Printf("程序执行结果是: %s\n", result)
	//fmt.Println("程序执行度过的时间是:", elapsed)
	r, _ := hex.DecodeString("ec01")
	fmt.Println(r)
	//s := time.Now().Format("2006-01-02 15:04:05")
	//fmt.Println(s, reflect.TypeOf(s))
}

func getSqrtSum(start float64, end float64) string {
	sum := 0.0
	if start>=end {
		return "null"
	}
	for i:=start; i<end; i++ {
		sum += math.Sqrt(i)
	}
	return fmt.Sprintf("%e", sum)
}
