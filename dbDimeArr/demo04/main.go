package main

import (
	"fmt"
)

func main() {
	var arr [3][4]int

	//1.输入数组
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("第%v组，第%v号的数字为：", i+1, j+1)
			fmt.Scanln(&arr[i][j])
		}
	}

	//2.首次遍历数组
	for i := 0; i < length; i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println(" ")
	}
	fmt.Println()
	//3.周围值清零
	for i := 0; i < length; i++ {
		for j := 0; j < len(arr[i]); j++ {
			if i == 0 || i == length-1 {
				arr[i][j] = 0
			} else if j == 0 || j == len(arr[i])-1 {
				arr[i][j] = 0
			}
		}
	}

	//4.二次遍历数组
	for i := 0; i < length; i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println(" ")
	}
}
