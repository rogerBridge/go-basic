package main

import "fmt"

// 定义学生这个类型
type student struct {
	name          string
	age           int // 大于23岁小于24岁, 按照24岁计算
	habits        []string
	otherProperty map[string]string
}

// 定义nr这个类型
type nr struct {
	l float64
	d float64
	r float64
}

func main() {
	//// 直接定义, 浪费内存
	//var b student
	//b.name = "leo"
	//b.age = 22
	//b.habits = make([]string, 0)
	//for i := 0; i < 100; i++ {
	//	b.habits = append(b.habits, string(i))
	//}
	//fmt.Println("结构体b所占用的字节数量是: ",unsafe.Sizeof(b))
	//
	//// new 创造的是: 指向结构体的指针, 声明结构体的方式1
	//a := new(student)
	//a.name = "安瑞峰"
	//a.age = 24
	//a.habits = []string{"王者荣耀", "杀手", "bilibili"}
	//a.otherProperty = map[string]string{"喜欢的人": "天天"}
	//// 毕竟, 指针在使用的时候确实非常的节省空间 :)
	//fmt.Println(a, unsafe.Sizeof(a))
	//
	//// 使用new定义nr
	//numberArray := new(nr)
	//fmt.Println(numberArray, unsafe.Sizeof(numberArray))
	//// 使用指针表达式定义nr
	//numberArray1 := &nr{1, 2, 3}
	//fmt.Println(numberArray1, unsafe.Sizeof(numberArray1))
	//// 实例化一个结构体
	////var numberArray2 nr = nr{4, 5, 6}
	//numberArray2 := nr{4, 5,6}
	//fmt.Println(numberArray2, unsafe.Sizeof(numberArray2))
	m := make(map[string]string)
	m["name"] = "leo2n"
	m["age"] = "25"
	type n struct {
		Name string
		Age  int
	}
	fmt.Println(len(m), m)
	nn := new(n)
	nn.Name = "leo2n"
	nn.Age = 25
	fmt.Println()

}
