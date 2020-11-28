package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	var p1 person
	p1.Name = "小王"
	fmt.Printf("p1的地址为%p\n", &p1)
	var p2 *person = &p1 //复制的操作
	//指针直接取值,需要打印地址必须使用%p格式化操作
	fmt.Printf("p2的地址为%p，p2储存的值为%p\n", &p2, p2)
	fmt.Println((*p2).Name)
	fmt.Println((*p2).Age)

	(*p2).Name = "小李"

	fmt.Printf("p2.name=%v,p1.name=%v", p2.Name, p1.Name)

}
