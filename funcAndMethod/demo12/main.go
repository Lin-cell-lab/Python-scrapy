package main

import (
	"fmt"
)

//将金字塔封装到函数内部
func printPyramid(totallevel int) {
	//使用for循环编写一个程序，可以接受一个整数，表示层数，打印一个金字塔

	for i := 1; i <= totallevel; i++ {
		//打印星号前打印空格

		for k := 1; k <= totallevel-i; k++ {
			fmt.Print(" ")
		}
		//j表示层数
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func printTable(num int) {

	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", i, j, i*j)
		}
		fmt.Println(" ")
	}
}

func main() {
	//调用函数打印金字塔
	var n int
	fmt.Println("请输入打印的层数：")
	fmt.Scanln(&n)
	printPyramid(n)

	var m int
	fmt.Println("请输入打印的层数：")
	fmt.Scanln(&m)
	printTable(m)

}
