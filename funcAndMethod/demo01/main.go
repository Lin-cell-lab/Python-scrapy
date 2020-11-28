package main

import (
	"fmt"
)

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test() n1=", n1)
}

//func 函数名 （形参列表）（返回值类型列表）{
//执行语句
// return返回值列表}
//1.支持返回多个返回值 2.希望忽略某个返回值，使用_符号表示占位忽略 3.返回值只有一个时，（返回值类型列表） 可以不写（）

func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getSum sum=", sum)
	return sum
}

//计算两个数的和与差，并返回结果
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

//函数的调用与回收机制，优化内存，提高效率
func main() {
	n1 := 10
	test(n1)
	fmt.Println("main() n1=", n1)
	sum := getSum(1, 2)
	fmt.Println("main() sum=", sum)
	//如果只想返回一个值，使用下划线占位
	//_, reslut2 := getSumAndSub(2, 4)
	result1, reslut2 := getSumAndSub(2, 4)
	fmt.Printf("sum=%v,sub=%v", result1, reslut2)
}
