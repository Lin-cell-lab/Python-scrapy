package main

import (
	"fmt"
)

//二维数组的应用案例，求班级平均分
func main() {

	var arr [3][5]float64

	for i := 0; i < len(arr); i++ {

		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("请输入第%d班第%d名同学的成绩：\n", i+1, j+1)
			fmt.Scanln(&arr[i][j])
		}

	}

	//遍历成绩
	sum := 0.0
	totalSum := 0.0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			sum += arr[i][j]
		}
		totalSum += sum
		fmt.Printf("第%d班的总分为%v,平均分为%v\n", i+1, sum, sum/float64(len(arr[i])))
	}
	fmt.Printf("所有班级的总分为%v,平均分为%.2f", totalSum, totalSum/15.0)
}
