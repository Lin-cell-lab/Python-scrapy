package main

import (
	"fmt"
)

func insertNum(arr []int, insertNum int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] >= insertNum {
			arr[i+1] = insertNum
			for k := i + 1; i < len(arr)-i-1; k++ {
				arr[k+1] = arr[k]
			}
		}

	}
}

//插入一个数字，数组顺序不变，打印数组
func main() {
	var arr []int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	fmt.Println("排序前arr=", arr)
	insertNum(arr, 3)

}
