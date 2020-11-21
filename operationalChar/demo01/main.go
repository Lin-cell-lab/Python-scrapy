package main

import (
	"fmt"
)

func main() {
	// /去小数留整数，加小数可保留运算后的小数（取模）
	fmt.Println(10 / 3)

	// % 取余数
	//公式：a%b = a-(a/b)*b
	fmt.Println(10 % 3)
	fmt.Println(10 - (10/3)*3)

	// ++和--的使用，a++和a--只能独立使用
	var i int = 10
	i++ // i=i+1
	fmt.Println("i=", i)
	i-- // i=i-1
	fmt.Println("i=", i)
}
