package main

import (
	"fmt"
)

func test01(arr *[3]int) {
	(*arr)[0] = 21
}

//数组细节
func main() {
	//1.数组是一个相同数据类型的组合，一个数组一旦声明了变量，其长度固定，不能动态变化
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	//arr[2] = 1.2 类型变化
	arr[2] = 12
	//arr[3] = 1.2 多余元素

	fmt.Println(arr)

	//2.数组中的元素可以是任何数据类型，包括值类型和引用类型，但是不可以混用

	//3.数组创建后，如果没有赋值，默认零值

	//4.越界错误 panic 下标过大

	//5.数组类型为值类型，默认值传递，因此会进行值拷贝，数组之间不会相互影响
	// 	arr1 := [3]int{12, 213, 231}
	// 	test01(arr1)
	// 	fmt.Println(arr1)
	arr1 := [3]int{92, 213, 231}
	test01(&arr1)
	fmt.Println("main arr1=", arr1)

	//5.长度是数组类型的一部分，在传递参数时需注意，考虑数组的长度。
}
