package main

import (
	"fmt"
)

func main() {
	var a byte
	var b int
	fmt.Println("请输入班级号和学号：")
	fmt.Scanf("%v , %d", &a, &b)
	fmt.Println(a, b)
}
