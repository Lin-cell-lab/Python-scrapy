package main

import (
	"fmt"
)

//1.作用域在整个包都有效，首字母大写在整个程序都有效(就近原则)

var age int = 30

var name string = "小米"

//Name ：= “tom”
//var Name string
//Name = “tom”不能执行：赋值语句不能在函数体外

func test01() {
	//局部变量只能在本函数内使用
	age := 1
	Name := "tom~"
	fmt.Println(age, Name)
}

func main() {
	test01()
	fmt.Println(age, name)

	//如果变量在一个代码块，比如for或者if循环，那么变量的作用域在该代码块

	for i := 0; i <= 10; i++ {
		fmt.Println("i=", i)
	}
	//	fmt.Println(i)
}
