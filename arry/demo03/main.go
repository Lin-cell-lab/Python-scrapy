package main

import (
	"fmt"
)

func main() {
	//for-range 遍历输出
	var heros [3]string = [3]string{"蜘蛛侠", "蝙蝠侠", "超人"}

	for index, value := range heros {
		fmt.Printf("index=%v,value=%v\n", index, value)
	}
}
