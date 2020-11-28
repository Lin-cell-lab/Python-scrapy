package main

import (
	"fmt"
)

func main() {
	//金字塔打印
	var totalSum float64 = 0.0
	var classNumber int = 1
	var studentNumber int = 2
	for j := 1; j <= classNumber; j++ { //j表示班级
		sum := 0.0
		for i := 1; i <= studentNumber; i++ {
			var score float64
			fmt.Printf("请输入第%v个班，第%v个同学的成绩:\n", j, i)
			fmt.Scanln(&score)
			sum += score
		}
		fmt.Printf("第%v班级的平均分是:%v\n", j, sum/float64(studentNumber))
		totalSum += sum
	}
	fmt.Printf("所有班级的总成绩为：%v,平均成绩为%v", totalSum, totalSum/(float64(classNumber)*float64(studentNumber)))
}
