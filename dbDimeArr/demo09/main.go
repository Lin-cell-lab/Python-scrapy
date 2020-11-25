package main

import (
	"fmt"
)

func findValue(arr [5]int) {
	maxVal := arr[0]
	maxValIndex := 0
	minVal := arr[0]
	minValIndex := 0

	for i := 0; i < len(arr); i++ {
		if maxVal < arr[i] {
			//最大值查找
			maxVal = arr[i]
			maxValIndex = i
		} else if minVal > arr[i] {
			//查找最小值
			minVal = arr[i]
			minValIndex = i
		}
	}
	fmt.Printf("最大值为%v，下标为%v\n最小值为%v，下表为%v", maxVal, maxValIndex, minVal, minValIndex)

}

func main() {
	var arr [5]int = [...]int{3, 45, 32, 1, 6}
	findValue(arr)
}
