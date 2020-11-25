package main

import (
	"fmt"
)

func main() {
	//常规遍历
	var arr [5]int = [...]int{1, 23, 3, 32, 32}
	slice := arr[1:4]
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	fmt.Println()
	//for-range遍历
	for _, value := range slice {
		fmt.Println(value)
	}

	//append函数追加切片
	var slice1 []int = []int{21, 213, 213, 23213}
	fmt.Println(slice1)
	//1.通过元素追加切片
	slice1 = append(slice1, 2121, 21309)
	fmt.Println(slice1)
	//2.通过切片追加切片
	slice1 = append(slice1, slice...)
	fmt.Println(slice1)

	//通过copy拷贝切片 ,两切片拷贝的容量大小比较，只拷贝仅有容量的元素，不报错
	var slice3 []int = []int{12, 21, 213, 12, 21}
	var slice4 = make([]int, 10)
	//slice3和slice4空间独立，相互不影响
	copy(slice4, slice3)
	slice3[2] = 333
	fmt.Println(slice4)
	fmt.Println(slice3)
}
