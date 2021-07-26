// Go 支持在结构体类型中定义_方法(methods)_ 。

package main

import (
	"fmt"
	"math"
)

type rect struct {
	width, height int
}

type circle struct {
	radius float64
}

// 计算○的周长
func (r *circle) perimeter() string {
	return fmt.Sprintf("%.2f", math.Pi*2*r.radius)
}
// 计算圆的面积
func (r *circle) area() float64 {
	return math.Pi*r.radius*r.radius
}

// 这里的 `area` 方法有一个_接收器(receiver)类型_ `rect`。
func (r *rect) area() int {
	return r.width * r.height
}

// 可以为值类型或者指针类型的接收器定义方法。这里是一个
// 值类型接收器的例子。
func (r *rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	//r := rect{width: 10, height: 5}
	c := circle{10}
	// 这里我们调用上面为结构体定义的两个方法。
	fmt.Println("area: ", c.area())
	fmt.Println("perim:", c.perimeter())

	// Go 自动处理方法调用时的值和指针之间的转化。你可以使
	// 用指针来调用方法来避免在方法调用时产生一个拷贝，或者
	// 让方法能够改变接受的结构体。
	rp := &c
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perimeter())
}

