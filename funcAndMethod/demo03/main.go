package main

import (
	"fmt"
)

//1.函数的作用域，局部变量
//2.基本数据类型都是默认值传递，即进行值拷贝，在函数内部修改，不会影响到原来的值
func test() {
	//n1为test的局部变量，只能在test中使用
	//var n1 int = 10
}

//数组的传递
//go不支持函数的重载
func test2(n1 int) {
	n1 = n1 + 10
	fmt.Println("test02() n1=", n1)
}

//希望函数内部的变量能够修改函数外部的变量，传入变量的地址，函数内以指针的方式操作变量，类似于引用
func test3(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("test3() n1=", *n1)
	fmt.Printf("n1的地址为%v\n", &n1)
}

func main() {
	n1 := 20
	test2(n1)
	fmt.Println("main() n1= ", n1)
	num := 20
	test3(&num)
	fmt.Println("main() num=", num)
	fmt.Printf("num的地址为%v\n", &num)
}
