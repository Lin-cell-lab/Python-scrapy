package main

import (
	"fmt"
	"unsafe"
)

//⽀持指针类型 *T，指针的指针 **T，以及包含包名前缀的 *<package>.T
//• 默认值 nil，没有 NULL 常量。
// • 操作符 "&" 取变量地址，"*" 透过指针访问⺫标对象。
// • 不⽀持指针运算，不⽀持 "->" 运算符，直接⽤ "." 访问⺫标成员。
// 不能对指针做加减法等运算。
// 可以在 unsafe.Pointer 和任意类型指针间进⾏转换。
func main() {
	var num int = 10
	fmt.Printf("num的地址是%v", &num)
	var ptr *int = &num
	*ptr = 12
	fmt.Printf("num的值为%d", num)

	x := 0x12345678
	p := unsafe.Pointer(&x) // *int -> Pointer
	n := (*[4]byte)(p)      // Pointer -> *[4]byte
	for i := 0; i < len(n); i++ {
		fmt.Printf("%X ", n[i])
	}

	// 将 Pointer 转换成 uintptr，可变相实现指针运算。

	d := struct {
		s string
		x int
	}{"abc", 100}
	p1 := uintptr(unsafe.Pointer(&d)) // *struct -> Pointer -> uintptr
	p1 += unsafe.Offsetof(d.x)        // uintptr + offset
	p2 := unsafe.Pointer(p)           // uintptr -> Pointer
	p1x := (*int)(p2)                 // Pointer -> *int
	*p1x = 200                        // d.x = 200
	fmt.Printf("%#v\n", d)
}
