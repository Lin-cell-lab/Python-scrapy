package main

import (
	"fmt"
)

func main() {
	// 1.100以内的数求和，求出当和第一次大于2哦的当前数
	var sum int = 0

	for i := 0; i <= 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Printf("%v的以前数字的和大于20\n", i)
			break
		}
	}

	//2.实现登录验证，有三次几会，用户名为"张无忌”，密码“999”,登录成功，否则提示还有几次输入机会
	var userName string = "张无忌"
	var passWord string = "999"
lable1:
	for t := 3; t > 0; t-- {
		fmt.Println("请输入用户名：")
		fmt.Scanln(&userName)
		fmt.Println("请输入密码：")
		fmt.Scanln(&passWord)

		if userName == "张无忌" && passWord == "999" {
			fmt.Println("输入正确！")
			break lable1
		} else {
			fmt.Printf("输入错误！您还剩【%v】次机会", t-1)
		}
	}
}
