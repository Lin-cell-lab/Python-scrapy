package main

import (
	"fmt"
)

//1.两个变量的交换 , 引用传递不修改原先变量
func swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Printf("a=%v , b=%v", a, b)
}
