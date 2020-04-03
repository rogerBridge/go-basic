package main

func main() {
	//isStructEqual()
	//outStudentInfo(inStudentInfo())

	//// 相比值结构体, 更常用指针结构体, 值传递和指针传递, 是否会作用到原来的结构体, 这个很关键
	//s := new(Student)
	//s.advantage = "before"
	//dealStruct(*s)
	//fmt.Println(s.advantage)
	//dealStructPoint(s)
	//fmt.Println(s.advantage)
	//
	//// methods with no struct
	//ss := str("hello")
	//ss.show()

	// interface 对 struct的抽象演示
	//playInterface()

	//// show interface assert
	//showAssert()
	//
	//// type switch
	//findType("leo2n")
	//s := make([]int, 10)
	//findType(s)
	//
	//// switch i.(type) 可以用来判断不同interface
	//findType2("leo2n")
	//findType2(&Asian{age: 24, name: "leo2n"})

	// goroutine
	playGoroutine()

}
