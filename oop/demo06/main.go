package main

import (
	"fmt"
)

//将结构体绑定到方法
type person struct {
	name string
}

func (p person) speak() {
	fmt.Printf("%v是一个好人！", p.name)
}

func (p person) test() {
	fmt.Println("test() name=", p.name)
}

func (p person) camulateSum() {

	res := 0
	for i := 0; i < 10000; i++ {
		res += i
	}
	fmt.Println(res)
}

func (p person) camulateSum1(n int) {

	res := 0
	for i := 1; i < n; i++ {
		res += i
	}
	fmt.Println(res)
}

func (p person) camulateSum2(n1 int, n2 int) int {
	//在通过一个变量去调用方法时，调用机制和函数一样，不一样的地方时，变量调用时，该变量也会作为一个参数传递到方法
	//（值类型值拷贝，引用类型地址拷贝）
	return n1 + n2
}

//方法是作用在指定数据类型上的，即（和指定的数据类型绑定），因此自定义数据类型都可以有方法，而不仅仅是struct
//test()方法和person类型绑定,test（）只能通过person类型的变量来调用，也不能使用其他的变量来调用
func main() {
	var person1 person
	person1.name = "小米"
	person1.test()
	var person2 person
	person2.name = "小邓"
	person2.speak()
	person2.camulateSum()
	person2.camulateSum1(213213)
	res2 := person2.camulateSum2(10, 1221)
	fmt.Println("res2=", res2)
}
