package main

import (
	"fmt"
)

func main() {
	map1 := make(map[int]int, 10)
	map1[25] = 44
	map1[11] = 52
	map1[42] = 71
	map1[32] = 51
	map1[10] = 53

	fmt.Println(map1)

}
