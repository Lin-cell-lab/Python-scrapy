package main

import "fmt"

func main() {
	x := 1
	//1取反
	fmt.Println(^x) // 0001  取反后为-0010

	//逻辑运算符的使用

	// 逻辑与&&(短路与)
	var age int = 49
	if age > 30 && age < 50 {
		fmt.Println("ok!")
	}
	if age > 30 && age < 40 {
		fmt.Println("ok~")
	}
	//逻辑或 || 短路或
	var age1 int = 49
	if age1 > 30 || age1 < 50 {
		fmt.Println("ok!")
	}
	if age1 > 30 || age1 < 40 {
		fmt.Println("ok~")
	}
	// 逻辑非 ！
	var age2 int = 49
	if age2 > 30 {
		fmt.Println("ok!")
	}
	if !(age2 > 30) {
		fmt.Println("ok~")
	}

}
