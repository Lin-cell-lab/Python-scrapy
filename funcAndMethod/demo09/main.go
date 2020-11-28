package main

import (
	"fmt"
)

//递归调用
// 1.斐波那契数
func fbn(n int) int {
	if n == 1 || n == 2 {
		n = 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
	return n
}

//练习题二
func cal(n int) int {
	if n == 1 {
		n = 3
	} else {
		n = 2*cal(n-1) + 1
	}
	return n
}

//猴子吃桃子
func peach(n int) int {
	if n == 10 {
		return 1
	}
	return ((peach(n+1) + 1) * 2)
}

func main() {
	fmt.Println(fbn(9))
	fmt.Println(fbn(3))
	fmt.Println(cal(1))
	fmt.Println(cal(5))
	fmt.Println("第一天的桃子数量=", peach(1))
}
