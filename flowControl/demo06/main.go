package main

import (
	"fmt"
)

func main() {
	var a int32
	var b int32
	fmt.Println("请输入a的值：")
	fmt.Scanln(&a)
	fmt.Println("请输入b的值：")
	fmt.Scanln(&b)

	c := a + b
	if c >= 50 {
		fmt.Println("hello，world！")
	}
}
