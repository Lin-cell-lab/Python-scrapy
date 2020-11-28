package main

import (
	"fmt"
)

func main() {
	//1.判断一个整数，属于哪个范围，大于小于0或者等于0
	var a int = 10

	if a > 0 {
		fmt.Printf("%v大于0\n", a)
	} else if a < 0 {
		fmt.Printf("%v小于0\n", a)
	} else {
		fmt.Printf("%v等于0\n", a)
	}

	//2.判断一个年份是否是润年
	var year int = 2021

	if (year%4 == 0) && (year%100 != 0) {
		fmt.Printf("%v是普通润年\n", year)
	} else if (year%100 == 0) && (year%400 == 0) {
		fmt.Printf("%v是世纪润年\n", year)
	} else {
		fmt.Printf("%v不是闰年！\n", year)
	}

	//3.判断一个数是否是水仙花数字 ，其各个数位上的数字立法和等于其本身，如153=1*1*1+3*3*3+5*5*5
	var b int = 154

	var c int = b % 10
	var d int = (b / 10) % 10
	var e int = b / 100

	if e*e*e+d*d*d+c*c*c == b {
		fmt.Printf("%v是水仙花数字\n", b)
	} else {
		fmt.Printf("%v不是水仙花数字\n", b)
	}

	//4.判断1000以内的水仙花数字

	/*	for f := 1; f <= 1000; f++ {

		}*/

	//5.保存用户和密码，判断用户是否为张三，密码是否为1234，进行提示！
	var userName string
	var pwd int

	fmt.Println("请输入用户名：")
	fmt.Scanln(&userName)
	fmt.Println("请输入密码：")
	fmt.Scanln(&pwd)

	if userName == "张三" && pwd == 1234 {
		fmt.Println("账户密码输入正确！")
	} else {
		fmt.Println("账户密码输入错误！")
	}

	//6.根据月份和年份，求出改月的天数（1-12）
}
