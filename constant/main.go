package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//1.常量值必须是编译期可确定的?*数字、字符串、布尔值?*
	const x, y int = 1, 2     // 多常量初始化
	const s = "Hello, World!" // 类型推断
	const (                   // 常量组
		a, b      = 10, 100
		c    bool = false
	)
	//2.在常量组中，如不提供类型和初始化值，那么视作与上⼀常量相同
	const (
		s1 = "abc"
		x1 // x = "abc"
	)
	fmt.Println(s1, x1)
	//3.常量值还可以是 len、cap、unsafe.Sizeof
	// 等编译期可确定结果的函数返回值。
	const (
		a1 = "abc"
		b1 = len(a1) //一个整数8字节
		c1 = unsafe.Sizeof(b1)
	)
	fmt.Println(c1)
	//4.关键字 iota 定义常量组中从 0 开始按⾏计数的⾃增枚举值。
	const (
		Sunday    = iota // 0
		Monday           // 1，通常省略后续⾏表达式。
		Tuesday          // 2
		Wednesday        // 3
		Thursday         // 4
		Friday           // 5
		Saturday         // 6
	)
	const (
		//TODO <<??

		_        = iota             // iota = 0
		KB int64 = 1 << (10 * iota) // iota = 1
		MB                          // 与 KB 表达式相同，但 iota = 2
		GB
		TB
	)
	fmt.Println(TB)

	//在同⼀常量组中，可以提供多个 iota，它们各⾃增⻓。
	const (
		A, B = iota, iota << 10 // 0, 0 << 10
		C, D                    // 1, 1 << 10
	)
	fmt.Println(A, B, C, D)
	//如果 iota ⾃增被打断，须显式恢复。
	const (
		A1 = iota // 0
		B1 // 1
		C1 = "c" // c 
		D1 // c，与上⼀⾏相同。
		E1 = iota // 4，显式恢复。注意计数包含了 C、D 两⾏。
		F1 // 5
	  )

	//可通过⾃定义类型来实现枚举类型限制。1-n为不同类型的字段
}
