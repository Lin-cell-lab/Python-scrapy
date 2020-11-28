package main

import (
	"fmt"
)

//string和slice的区别
func main() {
	//1.string底层为一个byte数组（因此string的内容不可以修改），可以进行切片处理提取内容
	str := "hello@xiaodeng"
	strSlice := str[6:]
	fmt.Println(strSlice)

	//2.字符串的修改
	slice2 := []byte(str)
	slice2[0] = 'Z'
	str = string(slice2)
	fmt.Println(str)

	//3.中文字符的修改[]rune(字符串)进行转码修改
	slice3 := []rune(str)
	slice3[0] = '北'
	str = string(slice3)
	fmt.Println(str)

}
