package main

import (
	"fmt"
)

func main() {
	//通过make方法创建切片,底层为一个数组
	var slice []float64 = make([]float64, 5, 10)
	slice[2] = 21
	slice[3] = 23
	fmt.Printf("slice=%v,slice的长度=%d，slice的容量为%v\n", slice, len(slice), cap(slice))

	fmt.Println()
	//直接定义和使用
	var strSlice []string = []string{"xiaoli", "xaiomi"}
	fmt.Println(strSlice)
	fmt.Println(len(strSlice))
	fmt.Println(cap(strSlice))
}
