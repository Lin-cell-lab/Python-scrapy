package main

import (
	"fmt"
)

//golang中的数据类型的转换
func main() {
	var i int = 100
	//1.将i->float
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	fmt.Printf("i=%v,n1=%v,n2=%v\n", i, n1, n2)
	//2.被转换的是变量储存的值即数据，变量本身并不会发生改变，（i的type仍然是int，只是把100转换）
	fmt.Printf("i's type is %T\n", i)
	//3.将int63转换成int32【-128-127】,编译时不报错，但是转换的结果按益出处理，转换时，考虑转换范围
	var num1 int64 = 99999
	var num2 int8 = int8(num1)
	fmt.Printf("num2's type is %T\n", num2)
	fmt.Println("num2=", num2)
}
