package main

import (
	"fmt"
)

func main() {
	//顺序查找
	var str [10]string = [...]string{"bc", "ac", "sa", "AA", "AA", "bc", "ac", "sa", "sa", "we"}

	var str1 string = "AA"
	index := -1
	for i := 0; i < len(str); i++ {
		if str1 == str[i] {
			index = i
			if index != -1 {
				fmt.Printf("找到了%v,下标为%v\n", str1, index)
			}
		}
	}
}
