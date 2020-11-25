package main

import (
	"fmt"
)

//二维数组
func main() {
	//1.打印方阵图像
	var arr [4][6]int
	//赋初始值
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	//遍历二维数组
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	//2.内存布局的分析
	var arr1 [2][3]int
	arr1[1][1] = 32
	fmt.Println(arr1)

	fmt.Printf("arr1[0]的内存地为%p,arr1[0][0]的内存地为%p\n", &arr1[0], &arr1[0][1])
	fmt.Printf("arr1[1]的内存地为%p,arr1[1][0]的内存地为%p\n", &arr1[1], &arr1[1][0])

	//3.
	var arr2 [2][3]int = [2][3]int{{1, 2, 3}, {1, 3, 2}}
	fmt.Println("arr2=", arr2)
}
