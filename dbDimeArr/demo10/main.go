package main

import (
	"fmt"
)

func insertNum(arr [5]int, insertNum int) {
	for i := 0; i < len(arr); i++ {
		if (arr[i] > insertNum) && (arr[i+1] < insertNum) {
			fmt.Println(arr[i], arr[i+1])
		}
	}
}

//插入一个数字，数组顺序不变，打印数组
func main() {
	var arr [5]int
	arr[0] = 6
	arr[1] = 4
	arr[2] = 1
	arr[3] = 4
	arr[4] = 9

	fmt.Println("排序前arr=", arr)
	insertNum(arr, 3)

}
