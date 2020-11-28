package main

import (
	"fmt"
)

func main() {
	//打压9的整数及和
	var max uint64 = 100
	var count uint64 = 0
	var sum uint64 = 0
	var i uint64 = 1

	for ; i <= max; i++ {
		if i%9 == 0 {
			count++
			sum += i
		}
	}
	fmt.Printf("count=%v,sum=%v\n", count, sum)
}
