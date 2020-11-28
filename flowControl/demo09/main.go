package main

import (
	"fmt"
)

func main() {
	// 1.大小写转换
	var char byte
	fmt.Println("请输入一个字母：")
	fmt.Scanf("%c", &char)

	switch char {
	case 'a':
		fmt.Println("A")
	case 'b':
		fmt.Println("B")
	case 'c':
		fmt.Println("C")
	case 'd':
		fmt.Println("D")
	case 'e':
		fmt.Println("E")
	default:
		fmt.Println("未找到相应匹配，请重新输入！")
	}

	// 2.成绩分档
	var score float64
	fmt.Println("请输入学生的成绩：")
	fmt.Scanln(&score)

	switch int(score / 60) {
	case 1:
		fmt.Println("及格！")
	case 0:
		fmt.Println("不及格！")
	default:
		fmt.Println("成绩输入有误，请重新输入！")
	}

	// 3.季节题目
	var month byte
	fmt.Println("请输入季节：")
	fmt.Scanln(&month)

	switch month {
	case 3, 4, 5:
		fmt.Println("春季")
	case 6, 7, 8:
		fmt.Println("夏季")
	case 9, 10, 11:
		fmt.Println("秋季")
	case 12, 1, 2:
		fmt.Println("冬季")
	default:
		fmt.Println("输入有误，未找到相应匹配！")
	}
}
