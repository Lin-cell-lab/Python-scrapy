package main

import (
	"fmt"
)

func main() {
	//引⽤类型包括 slice、map 和 channel。
	//1.定义完后 默认值为0,
	var intArr [3]int
	fmt.Println(intArr)

	//2.数组的地址为数组第一个元素地址,后一个 字节为前面一个地址加8（int32为4个字节，int64为8个字节）
	fmt.Printf("intarr的地址为%p,intarr[0]的地址为%p,intarr[1]的地址为%p,intarr[2]的地址为%p",
		&intArr, &intArr[0], &intArr[1], &intArr[2])

	//练习题1，数组的打印
	var score [5]float64

	for i := 0; i < len(score); i++ {
		fmt.Println("请输入具体的数值：\n", i+1)
		fmt.Scanln(&score[i])
	}
	//变量数组打印
	for i := 0; i < len(score); i++ {
		fmt.Printf("score[%d]=%v\n", i, score[i])
	}

	//四种初始化的方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01=\n", numArr01)

	var numArr02 = [3]int{1, 2, 3}
	fmt.Println("numArr02=\n", numArr02)

	var numArr03 = [...]int16{8, 2, 3}
	fmt.Println("numArr03=\n", numArr03)

	var numArr04 = [...]int{1: 8, 0: 2, 2: 3}
	fmt.Println("numArr04=\n", numArr04)

	var strArr05 = [...]string{1: "小李", 2: "小红", 0: "小白"}
	fmt.Println("strArr05=\n", strArr05)
}
