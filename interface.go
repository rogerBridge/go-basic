package main

import (
	"fmt"
	"math"
)

// 接口就是把一系列类的通用方法抽象出来
// 比如: 圆和正方体都有方法: 面积, 周长

type Geometry interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

func (r *Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * r.Length * r.Width
}

// interface 就是对struct的共有属性的一些抽象
func play1() {
	c1 := new(Circle)
	c1.Radius = 1
	r1 := new(Rectangle)
	r1.Width = 1
	r1.Length = 2
	a := []Geometry{c1, r1}
	for _, v := range a {
		fmt.Println(v.Area(), v.Perimeter())
	}
}

// type assert
// 采用: v, ok := i.(int), 可以根据ok的true||false来判断是否需要执行下一步
func assert(i interface{}) {
	v, ok := i.(int)
	if ok == true {
		fmt.Println(v)
	} else {
		fmt.Println(i, "can't get type assert")
	}
}

func showAssert() {
	var s interface{} = "steve jobs"
	assert(s)
	var i interface{} = 20
	assert(i)
}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("I am a string", i.(string))
	case int:
		fmt.Println("I am a int", i.(int))
	case float64:
		fmt.Println("I am a float64", i.(float64))
	default:
		fmt.Println("Other type :)")
	}
}

// 肿么理解"万物皆是接口"呢?
type Person interface {
	//Dress()
	//Eat()
	Sleep()
	//Run()
}

type Asian struct {
	name string
	age  int
}

func (p *Asian) Sleep() {
	fmt.Printf("A person name %s and age %d is sleeping", p.name, p.age)
}

func findType2(i interface{}) {
	switch i.(type) {
	case Person:
		fmt.Println("type is Person")
	default:
		fmt.Println("unknown type")
	}
}
