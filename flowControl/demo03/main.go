package main

import (
	"fmt"
)

func main() {
	var score float32
	fmt.Println("请输入成绩：")
	fmt.Scanln(&score)

	if score < 8.0 {
		var gender string
		fmt.Println("恭喜进入决赛！请输入您的性别(male or female)：")
		fmt.Scanln(&gender)

		if gender == "male" {
			fmt.Println("请进入男子组！")
		} else {
			fmt.Println("请进入子女组")
		}

	} else {
		fmt.Println("很遗憾，您未进入决赛！")
	}
}
