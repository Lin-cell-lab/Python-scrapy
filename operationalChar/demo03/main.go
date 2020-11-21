package main

import (
	"fmt"
)

func main() {
	//放假例题
	var allDays int = 97
	var allWeeks int = allDays / 7
	var lastDays int = allDays % 7
	fmt.Printf("weekds=%d,days=%d\n", allWeeks, lastDays)
	//温度例题
	var hausi float32 = 123.23
	var sheshi float32 = 5.0 / 9 * (hausi - 100) //记得小数的使用
	fmt.Printf("%v华氏度对应的摄氏度为%v", hausi, sheshi)

}
