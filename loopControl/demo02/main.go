package main

import (
	"fmt"
)

// 1.循环变量初始化 2.循环条件 3.循环体 4. 循环变量迭代
func main() {
	//1.循环条件返回一个布尔值

	for i := 1; i <= 10; i++ {
		fmt.Println("hello，world", i)
	}

	k := 2
	for k <= 90 {
		fmt.Println("hello，world", k)
		k++
	}
}
