package main

import (
	"fmt"
)

//3.全局匿名函数
var (
	//func1代表一个全局匿名函数
	func1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	//匿名函数1.在定义匿名函数时直接调用，只能调用一次。
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	//2.将匿名函数赋值个一个变量，则a的类型就是函数类型，通过a则可以调用该匿名函数
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	res1 := a(10, 30)

	res2 := func1(10, 1212)

	fmt.Println("res=", res, "res1=", res1, "res2=", res2)
}
