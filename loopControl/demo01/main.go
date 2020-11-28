package main

import (
	"fmt"
)

func main() {
	//传统遍历字符串
	var str string = "hello！北京"
	str3 := []rune(str)
	for i := 0; i < len(str3); i++ {
		fmt.Printf("%c \n", str3[i])
	}

	//采用for-range,如果字符串含有中文，那么传统的字符串遍历会出现乱码，utf-8对应3个字节，推荐使用，默认使用字符
	var str2 string = "lakers finals champion！北京"
	for index, val := range str2 {
		fmt.Printf("index=%x ,val=%c \n", index, val)
	}
}
