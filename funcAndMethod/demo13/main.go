package main

import (
	"fmt"
)

//golang中的方法作用在指定的数据类型上，自定义数据类型都可以有方法

type integer int

func (i integer) print() {
	fmt.Println("i = ", i)
}

//指针修改i的值 *取i的值（开），&取地址（关）
func (i *integer) change() {
	*i = *i + 1
	fmt.Println(*i)
}

// 如果一个类型实现了string方法，则定义变量
type students struct {
	name string
	age  int
}

//students 的方法
func (stu *students) string() string {
	str := fmt.Sprintf("name=[%v] age=[%v]", stu.name, stu.age)
	return str
}

func main() {
	var i integer = 30
	i.print()
	i.change()

	//如果实现了*student的方法，就会自动调用
	str1 := students{
		name: "小米",
		age:  12,
	}
	fmt.Println(str1)
}
