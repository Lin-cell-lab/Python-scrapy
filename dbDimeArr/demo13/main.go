package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//1.随机生成十个整数(1-100)
	//2.并倒序打印
	//3.求平均值，最大值，和最大值的下标
	//4.并查找里面是否有55

	//生成随机数
	rand.Seed(time.Now().UnixNano()) //将种子放在循环外面

	var arr1 [10]int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = rand.Intn(100)
	}

	//倒序打印
	fmt.Println("正=", arr1)

	length := len(arr1)
	temp := 0
	for i := 0; i < length/2; i++ {
		temp = arr1[length-1-i]
		arr1[length-1-i] = arr1[i]
		arr1[i] = temp
	}
	fmt.Println("反=", arr1)

	//求平均值
	sum := 0.0
	for i := 0; i < length; i++ {
		sum += float64(arr1[i])

	}

	//找最大值，并得出下标
	maxValue := arr1[0]
	maxValueIndex := 0

	for i := 0; i < length; i++ {
		if maxValue < arr1[i] {
			maxValue = arr1[i]
			maxValueIndex = i
		}
	}
	fmt.Printf("最大值为%v,下标为%v\n", maxValue, maxValueIndex)
	//找存在值
	index := -1
	findValue := 55
	for i := 0; i < length; i++ {
		if findValue == arr1[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("已经找到%v\n", findValue)
	} else {
		fmt.Printf("未找到%v\n", findValue)
	}

	fmt.Printf("总和为%v,平均值为%v,", sum, sum/float64(length))
}
