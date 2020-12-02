package main

import (
	"fmt"
)

type aInterface interface {
	Say()
}

type students struct {
	Name string
}

func (stu students) Say() {
	fmt.Println("stu say()...")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i =", i)
}

type bInterface interface {
	Hello()
}

type monster struct {
}

func (m monster) Hello() {
	fmt.Println("monster Hello()")
}

type t interface {
}

func main() {
	//1.接口本身不能创建实例，但可以指向一个实现了该接口的自定义类型的变量（实例）
	var stu students //结构体变量实现了say（）方法，实现了ainterface
	var a aInterface = stu
	a.Say()
	//2.接口所有的方法都没有结构体，即都是没有实现的方法
	//3.一个自定义类型需要将某个接口的所有方法都实现，我们说这个自定义类型实现了该接口
	//4.一个自定义类型只有实现了某个接口，才能将该自定义类型的实例赋给接口类型
	//5.只要是自定义类型数据，就可以实现接口，不仅仅是结构体变量
	//6.一个自定义类型可以实现多个接口
	var i integer = 10
	var b aInterface = i
	b.Say()

	//7.接口里面不能有变量
	//8.接口之间可以相互继承
	//9.interface为一个引用类型，指针，空值为nil
	//10.空接interfa{}无任何方法，任何类型都实现了空接口,任何变量都可以赋给空接口
	var t t = stu
	fmt.Println(t)
	var t2 interface{} = stu
	var num1 float64 = 8.8
	t2 = num1
	fmt.Println(t2, t)
}
