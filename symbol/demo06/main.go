package main

import "fmt"

func main() {
	// 演示goto的使用
	goto lable1
	fmt.Println("ok1")
	fmt.Println("ok2")
	fmt.Println("ok3")
lable1:
	fmt.Println("ok4")
}
