package main

import (
	"fmt"
)

func main() {
	var score int
	fmt.Println("请输入成绩：")
	fmt.Scanln(&score)

	if score == 100 {
		fmt.Println("奖励一台BMW")
	} else if score > 80 && score <= 99 {
		fmt.Println("奖励一台苹果七")
	} else if score >= 60 && score <= 80 {
		fmt.Println("奖励一台ipaid")
	} else {
		fmt.Println("滚蛋垃圾！")
	}
	//找到一个接口就会执行，不会同时执行
}
