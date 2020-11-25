package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr [4][4]int
	//1.随机数生成二维数组
	rand.Seed(time.Now().UnixNano()) //将种子放在循环外面

	for i := 0; i < len(arr); i++ {

		for j := 0; j < len(arr[i]); j++ {

			arr[i][j] = rand.Intn(10)
		}
	}

	// //1.获取值
	// for i := 0; i < len(arr); i++ {
	// 	for j := 0; j < len(arr[i]); j++ {
	// 		fmt.Printf("第%v组，第%v号的数字为：", i+1, j+1)
	// 		fmt.Scanln(&arr[i][j])
	// 	}
	// }

	//2.首次遍历数组
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println(" ")
	}

	fmt.Println()

	//两组交换
	arr[0], arr[3] = arr[3], arr[0]
	arr[1], arr[2] = arr[2], arr[1]

	//2.再次打印数组
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println(" ")
	}
}
