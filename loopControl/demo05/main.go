package main

import (
	"fmt"
)

func main() {
	//使用for循环编写一个程序，可以接受一个整数，表示层数，打印一个金字塔
	var totallevel int = 11
	for i := 1; i <= totallevel; i++ {
		//打印星号前打印空格

		for k := 1; k <= totallevel-i; k++ {
			fmt.Print(" ")
		}
		//j表示每层个数
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
