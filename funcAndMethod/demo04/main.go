package main

import (
	"fmt"
)

//1.在go中，函数也是一种数据类型，可以赋值给一个变量，改变量就是一个函数类型的变量，通过该变量可以对函数调用
func sum(n1 int, n2 int) int {
	return n1 + n2
}

//4.此时的myfun就是一个func（int,int） int 类型
type myFuncType func(int, int) int

//2.函数作为一种数据类型，在go中，可以作为形参传递
func myFun(funvar myFuncType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

//3.自定义数据类型
func test4() {
	type myInt int
	var num1 myInt = 30
	// var num2 int
	// num2 = num1 区别类型不同  要用 num2 = int（num1）
	fmt.Printf("num1的数据类型是%T,，num=%v", num1, num1)
}

//5.支持对返回值命名
func test5(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

//6.编写一个getsum函数，求出1到多个int的和,可变参数的使用，放在形参列表的最后面
func sumGet(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i] //args【0】表示切片
	}
	return sum
}

func main() {
	a := sum
	//指向同一块的空间
	fmt.Printf("a的类型是%T，sum的类型是%T\n", a, sum)

	res := a(10, 23)
	fmt.Println(res)
	res2 := myFun(sum, 50, 60)
	fmt.Println("res2=\n", res2)

	test4()

	res3 := myFun(sum, 500, 1000)
	fmt.Println("res3=", res3)

	sum, sub := test5(90, 20)
	fmt.Println(sum, sub)

	res4 := sumGet(10, 21, 90)
	fmt.Println("res4=", res4)
}
