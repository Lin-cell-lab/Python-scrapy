package main

import (
	"fmt"
)

func main() {
	var month byte
	var age byte

	fmt.Println("请输入月份（1-12）：")
	fmt.Scanln(&month)
	fmt.Println("请输入您的年龄：")
	fmt.Scanln(&age)

	if month >= 4 && month <= 10 {
		if age <= 18 {
			fmt.Println("票价：30")
		} else if age > 18 && age < 60 {
			fmt.Println("票价：60￥")
		} else {
			fmt.Println("票价：20￥")
		}
	} else {
		if age >= 18 {
			fmt.Println("票价：40￥")
		} else {
			fmt.Println("票价：20￥")
		}
	}
}
