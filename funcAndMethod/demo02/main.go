package main

import (
	"fmt"
)

//递归调用
func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n=", n)
}

func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("n=", n)
	}

}

//函数的递归调用
func main() {
	test2(4)
}
