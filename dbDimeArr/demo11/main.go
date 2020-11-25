package main

import (
	"fmt"
	"math/rand"
	"time"
)

//保存奇数到数组并倒序打印
func main() {
	rand.Seed(time.Now().UnixNano())

	var arr [5]int
	//1.随机生成奇数组
	for i := 0; i < len(arr); i++ {
		num := rand.Intn(10)
		if num/2 != 0 {
			arr[i] = num
		}
	}
	fmt.Println(arr)

	temp := 0

	for i := 0; i < len(arr)/2; i++ {
		temp = arr[len(arr)-1-i]
		arr[len(arr)-1-i] = arr[i]
		arr[i] = temp
	}
	//2.打印数组
	fmt.Println(arr)
}
