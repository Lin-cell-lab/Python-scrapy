package main

import (
	"fmt"
)

//案例1，面积返回
type circle struct {
	radius float64
}

func (s circle) area() float64 {
	return s.radius * s.radius * 3.12
}

//案例2.为了提高传输效率，方法和结构体的指针类型进行绑定
func (c *circle) area2() float64 {
	//因为c为指针，标准访问方式为（*c）.redius
	//建议写出*c，
	c.radius = 10.0
	fmt.Printf("area2()中c指向的地址为%p\n", c)
	fmt.Printf("area2()中c本身的地址为%p\n", &c)

	return 3.14 * (*c).radius * (*c).radius

}

func main() {
	var square circle
	square.radius = 21.2
	s1 := square.area()
	fmt.Println("s1=", s1)

	var c circle
	(&c).radius = 5.0
	fmt.Printf("main()中c的地址为%p\n", &c)
	res2 := (&c).area2()
	fmt.Println("res2的面积=", res2)
	fmt.Println((&c).radius)
}
