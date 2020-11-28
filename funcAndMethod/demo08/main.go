package main

import (
	"fmt"
)

//defer函数
//当执行到defer时,暂时不执行,会将defer后面的语句压入到独立的栈(defer栈)
//当函数执行完毕以后,在defer栈中先入后出的方式出栈
func sum(n1 int, n2 int) int {
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)
	//在defer将语句放入栈时,也会将相应的值同时拷贝到栈
	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}
func main() {
	res1 := sum(10, 20)

	fmt.Println("res1=", res1)
}
