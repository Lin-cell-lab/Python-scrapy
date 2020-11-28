package main

import (
	"fmt"
)

func main() {
	var a, b int32 = 2, 13
	c := a + b

	if c%3 == 0 && c%5 == 0 {
		fmt.Printf("%v和%v的既能被3整除，也能被5整除。", a, b)
	}
}
