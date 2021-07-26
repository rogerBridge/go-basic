package main
import "fmt"

func main()  {
	// int32 ~ rune, int8 ~ byte
	//s := "hello, 安瑞峰"
	//c := []byte(s)
	//r := []rune(s)
	//i := []int32(s)
	//fmt.Println(c, r, i)

	q := make([]int, 10)
	for i:=0; i<100; i++{
		q = append(q, i)
	}
	fmt.Println(q)
}
