package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	//创建
	p2 := person{"marry", 21}
	fmt.Println(p2)

	//2.引用创建
	var p3 *person = new(person)
	(*p3).Name = "Smith"
	p3.Name = "john" //底层优化
	(*p3).Age = 21
	fmt.Println(*p3)

	//3.引用创建2
	var p4 *person = &person{"marry", 21}
	(*p4).Name = "小王"
	fmt.Println(*p4)
}
