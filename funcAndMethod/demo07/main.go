package main

import (
	"fmt"
	"strconv"
)

//闭包：一个函数与其相关的引用环境组合的一个整体（实体）
//1.addupper是一个函数，返回的数据类型为func（int） int
//2.返回的是一个匿名函数，但是匿名函数引用了函数外的n，因此这个匿名函数和n就形成了一个整体，构成闭包
//3.闭包是类，函数是操作，n为字段
//4.反复调用f函数时，n只初始化一次，因此没调用一次就进行累加
//5.闭包的关键就时要分析出返回的函数和他所使用（引用的变量），因为函数和其引用到的变量构成闭包
func addUpper() func(int) int {
	//累加器
	//初始化n和str，不调用
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += "abc"
		fmt.Println("str=", str)
		return n
	}

}

func strAdd() func(string, int, int) string {

	fmt.Println("第一层")
	return func(str string, num1 int, num2 int) string {
		n3 := num1 * num2
		str = str + strconv.Itoa(n3)
		return str
	}
}

func main() {
	//返回的就是一个函数k，将k当作函数使用
	k := strAdd()
	l := k("小米", 20, 29)
	fmt.Println(l)

	f := addUpper()
	add := func(base int) func(int) int {
		return func(n int) int {
			return base + n
		}
	}

	add5 := add(5)
	fmt.Println(add5(20))

	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
