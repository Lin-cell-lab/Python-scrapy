package main

import (
	"fmt"
)

type methodUtils struct {
	//一个空的结构体
}

//1.给methodutils编写方法打印10*8的*
func (mu methodUtils) print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println(" ")
	}
}

//2.提供m和n两个参数，打印m*n个矩形
func (mu methodUtils) ptint1(n int, m int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println(" ")
	}
}

//3.打印面积
func (mu methodUtils) eara(len float64, width float64) float64 {
	return len * width
}

//4.判断一个数的奇偶性
func (mu *methodUtils) num(num int) {
	if (num % 2) == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
}

//5.根据行类打印对应的字符
func (mu *methodUtils) print3(n int, m int, key string) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(key)
		}
		fmt.Println(" ")
	}
}

//定义caculator结构，实现加减乘除的功能
type caculator struct {
	num1 float64
	num2 float64
}

func (cau *caculator) getvalue(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = cau.num1 + cau.num2
	case '-':
		res = cau.num1 - cau.num2
	case '*':
		res = cau.num1 * cau.num2
	case '/':
		res = cau.num1 / cau.num2
	default:
		fmt.Println("符号输入有误！")
	}
	return res
}

func main() {
	var mu methodUtils
	mu.print()
	fmt.Println("")
	var mu1 methodUtils
	mu1.ptint1(10, 21)

	var mu2 methodUtils
	aera1 := mu2.eara(1.0, 12.2)
	fmt.Println("eara1 = ", aera1)

	(&mu).num(12)

	mu2.print3(2, 3, "&")

	var caculator1 caculator
	caculator1.num2 = 2.2
	caculator1.num2 = 80.7
	res1 := caculator1.getvalue('*')
	fmt.Println("res1 = ", res1)
	/* TODO */
	//计算错误待修改
}
