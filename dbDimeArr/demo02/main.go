package main

import (
	"fmt"
)

func main() {
	//1.二维数组的普通遍历
	var arr2 [2][3]int = [2][3]int{{1, 2, 3}, {1, 3, 2}}

	for i := 0; i < len(arr2); i++ {

		for j := 0; j < len(arr2[i]); j++ {

			fmt.Printf("%v\t", arr2[i][j])
		}
		fmt.Println()
	}
	//2.for-range遍历二维数组
	for i, v := range arr2 {
		fmt.Printf("i=%v,v=%v", i, v)
		for j, v2 := range v {
			fmt.Printf("j=%v,v2=%v", j, v2)
		}
		fmt.Println()
	}
}
