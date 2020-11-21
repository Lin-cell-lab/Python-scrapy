package main

import (
	"fmt"
)

func main() {
	//基本数据类型转换为string
	var num1 int = 99
	var num2 float64 = 23.222
	var b bool = true
	var myChar byte = 'h'
	var str string // 空的string

	//Sprintf根据format参数生成格式化的字符串并返回该字符串。
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str's type is %T   str=%v\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str's type is %T  str=%v\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str's type is %T str=%q\n", str, str)
	
	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str's type is %T str=%q", str, str)

}
