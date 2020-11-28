package main

import (
	"fmt"
)

//结构体的字段类型是：slice，指针，和map的零值都是nil，即还没有分配空间
//如果要先使用，则需要先make

type person struct {
	Name   string
	Age    int
	scores [5]int
	ptr    *int
	slice  []int
	map1   map[string]string
}

type monsters struct {
	Name   string
	Age    int
	scores [5]int
	ptr    *int
	slice  []int
	map1   map[string]string
}

func main() {
	var p1 person
	fmt.Println(p1)

	if p1.ptr == nil {
		fmt.Println("ok1")
	}

	if p1.map1 == nil {
		fmt.Println("ok2")
	}

	if p1.slice == nil {
		fmt.Println("ok3")
	}

	//1.make切片
	p1.slice = make([]int, 10)
	p1.slice[0] = 21
	fmt.Println(p1.slice)

	//2.makemap
	p1.map1 = make(map[string]string)
	p1.map1["小米"] = "笑"
	fmt.Println(p1.map1)

	//结构体为值类型
	var monster1 monsters
	monster1.Name = "牛魔王"
	monster1.Age = 211221
	fmt.Println(monster1)
}
