package main

import (
	"fmt"
)

func main() {
	//1.要求打印1-100之间的奇数，要求使用for循环+continue
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("基数i=", i)
	}

	//2.从键盘读入个数不确定的整数，并判断读入正数和负数的个数，输入为0时结束程序
	var positiveCount int
	var negativeCount int
	var num int
	for {
		fmt.Println("请输入一个整数：")
		fmt.Scanln(&num)

		if num == 0 {
			break //终止循环
		}

		if num > 0 {
			positiveCount++
			continue //结束本次循环进入下次
		}
		negativeCount++
	}
	fmt.Printf("正数有%v个，负数有%v个\n", positiveCount, negativeCount)

	//3.交现金路口费用
	var cash float64 = 100000.0
	var count int = 0
	for {
		if cash > 50000.0 {
			cash = cash - cash*0.05
			fmt.Printf("此次应该交213%v\n", cash)
			count++
		} else if cash < 50000.0 && cash > 0 {
			cash = cash - 1000
			fmt.Printf("此次应当交%v\n", cash)
			count++
		} else if cash <= 0 {
			break
		} else {
			fmt.Println("")
		}
	}
	fmt.Println(count)
}
