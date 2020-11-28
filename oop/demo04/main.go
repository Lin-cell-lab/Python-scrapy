package main

import (
	"fmt"
)

//结构体的镶套
type piont struct {
	x int
	y int
}

type rect struct {
	leftUp, rightDown *piont
}

func main() {
	//结构体所有的字段在内存中是连续的，指针本生的地址是连续的，但指针指向的地址不一定是连续的
	r1 := rect{&piont{1, 2}, &piont{2, 5}}
	//字节的连续分布：依靠地址值的加减寻址
	fmt.Printf("r1.leftUp.x:%p\nr1.leftUp.y:%p\nr1.rightDwon.x:%p\nr1.rightDwon.y.:%p",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)
	r2 := rect{&piont{12, 32}, &piont{34, 21}}
	fmt.Printf("r2.leftUp:%p\nr2.leftUp:%p\nr2.rightDwon:%p\nr2.rightDwon:%p",
		&r2.leftUp, &r2.leftUp, &r2.rightDown, &r2.rightDown)
}
