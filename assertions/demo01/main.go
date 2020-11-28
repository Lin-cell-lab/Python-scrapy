package main

import (
	"fmt"
)

type point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point1 point = point{1, 2}
	a = point1 // 将一个a赋给piont变量？
	var b point
	b = a.(point) //类型断言，判断a是否是指向piont类型的变量，如果是就转换成piont类型并赋给b变量
	fmt.Println(b)
	fmt.Println("only to we have sufficient contact and trust each other,",
		"we can rise above all the troubles and triumph!")

	//2.其他案例（待检测的类型断言）
	var x interface{}
	var b1 float32 = 1.2
	x = b1 //x转换为float32

	if y, ok := x.(float32); ok {
		fmt.Printf("y的类型为%T，值为%v", y, y)
		fmt.Println("转换成功")
	} else {
		fmt.Println("转换失败")
	}
	fmt.Println("继续执行！")
}
