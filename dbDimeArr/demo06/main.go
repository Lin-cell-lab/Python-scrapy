package main

import (
	"fmt"
)

func main() {
	var arr [8]int = [...]int{1, 23, 21, 43, 65, 23, 8, 3}

	averVale := 0.0
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	averVale = float64(sum) / float64(len(arr))
	bigAccount := 0
	SmallAccount := 0

	for i := 0; i < len(arr); i++ {
		if float64(arr[i]) > averVale {
			bigAccount++
		} else if float64(arr[i]) < averVale {
			SmallAccount++
		} else {
			return
		}
	}
	fmt.Println("平均值=", averVale)
	fmt.Printf("大于平均值的数有%v个\n小于平均值的有%v个", bigAccount, SmallAccount)
}
