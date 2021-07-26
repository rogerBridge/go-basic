package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64{
	return math.Pi*c.radius*c.radius
}

func (c *Circle) Perimeter() float64  {
	return 2*math.Pi*c.radius
}

type Rectangle struct {
	length float64
	width  float64
}

func (r *Rectangle) Area() float64 {
	fmt.Println(r.length*r.width)
	fmt.Printf("长:%f, 宽:%f 的矩形面积是:%f\n", r.length, r.width, r.width*r.length)
	return r.length*r.width
}

func (r *Rectangle) Perimeter() float64 {
	fmt.Printf("长:%f, 宽:%f 的矩形周长是:%f\n", r.length, r.width, 2*r.length*r.width)
	return 2*r.length*r.width
}

// 集合很像鸭子类型
type Geometry interface {
	Area() float64
	Perimeter() float64
}

func main() {
	r := &Rectangle{3, 5}
	c := &Circle{4}
	geo := []Geometry{r, c}
	for _, v := range geo {
		fmt.Println(v.Area(), v.Perimeter())
	}
	//d := Shaper(r)
	//fmt.Println(reflect.TypeOf(d))
}