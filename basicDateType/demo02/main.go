package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b = false
	fmt.Println("b=", b)
	// 注意事项
	//1.bool类型占用的储存空间是一个字节
	fmt.Println("b的占用空间是=", unsafe.Sizeof(b)) //1
	//2.bool类型只能取true或者false ，不以0或者非0指数来代替
}
