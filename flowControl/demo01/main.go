package main

import (
	"fmt"
)

func main() {
	var age int
	fmt.Println("请输入你爹年龄：")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Println("你的年龄大于18岁！")
	} else {
		fmt.Println("你的年龄小于18岁！")
	}
	// 	if age > 18 {
	// 		fmt.Println("你的年龄大于18岁！
}
