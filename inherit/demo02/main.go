package main

import (
	"fmt"
)

type a struct {
	Name string
	age  int
}

type b struct {
	Name  string
	score float64
}

type c struct {
	a
	b
	//Name string
}

type brand struct {
	Name    string
	Address string
}

type goods struct {
	Name  string
	Price float64
}

type tv struct {
	*goods
	*brand
}

func main() {
	//1.如果c无匿名结构体字段，而a和b都要name，这时就必须通过指定匿名结构体来区分
	var c1 c
	c1.a.Name = "小米"
	fmt.Println(c1.a.Name)

	//2.嵌套结构体之后，可以在创建结构体变量的同时，指定各个匿名结构体的字段的值
	// tv1 := tv{goods{"手机", 2131.2}, brand{"小米", "成都"}}
	// fmt.Println(tv1)

	tv2 := tv{&goods{"手机", 2131.2}, &brand{"华为", "北京"}}
	fmt.Println("tv3", *tv2.goods, *tv2.brand)
}
