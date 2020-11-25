package main

import (
	"fmt"
)

func main() {
	var intArr [8]int = [...]int{1, 4, 9, 9, 2, 3, 23, 23}
	slice := intArr[1:4]
	//切片的最大容量为数组下标的最大值
	fmt.Printf("slice = %v\nslice的容量 = %v\nslice的长度 = %d\n", slice, cap(slice), len(slice))
	//slice为一个引用类型，包含第一元素的地址，和长度，和容量(为一个structure结构体)
	fmt.Printf("slice[0]的地址为%p\n", &slice[0])
	fmt.Printf("intArr[1]的地址为%p\n", &intArr[1])
	fmt.Printf("slice的地址为%p\n", &slice)
	slice[1] = 231
	fmt.Println(slice)

}
