package main

import (
	"fmt"
)

type box struct {
	length float64
	width  float64
	hight  float64
}

func (cube *box) stere() float64 {
	return cube.hight * cube.length * cube.width
}

func main() {
	var cube1 = box{21.2,
		21.1,
		12.1}

	fmt.Println(cube1.stere())
}
