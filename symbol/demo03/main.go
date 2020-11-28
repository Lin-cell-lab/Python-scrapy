package main

import (
	"fmt"
)

//brake出现在多层镶套的语句块时，通过标签指明要终止的那一个一句块
func main() {
	//设置一个标签 lable2
lable2:
	for i := 0; i < 4; i++ {

		for j := 0; j < 10; j++ {
			if j == 2 {
				// brake 默认跳出最近的for 循环

				break lable2 // 指定标签，跳出对应的for循环
			}
			fmt.Println("j=", j)
		}
	}
}
