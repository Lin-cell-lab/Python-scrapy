package main

import (
	"fmt"
)

func main() {

	//数组的基本使用

	//1.定义数组
	var hens [6]float64
	//2.给数组元素赋值
	hens[0] = 3.0
	hens[1] = 4.2
	hens[2] = 3.0
	hens[3] = 4.1
	hens[4] = 3.0
	hens[5] = 3.9
	//3.遍历数组求出总体重
	totalweight := 0.0
	for i := 0; i < len(hens); i++ {
		totalweight += hens[i]
	}
	//求出凭借体重
	averagetotal := fmt.Sprintf("%.3f", totalweight/float64(len(hens)))
	fmt.Printf("totalweight=%.2f, averagetotal= %v", totalweight, averagetotal)

}
