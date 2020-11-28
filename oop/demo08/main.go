package main

import (
	"fmt"
)

type person struct {
	name string
}

//函数
//对于普通函数，接收者为值类型，不能将指针类型的数据直接传递
func test01(p person) { //值拷贝
	fmt.Println(p.name)
}

func test02(p *person) {
	fmt.Println((*p).name)
	//地址拷贝
}

//对于方法（struct）
//接受者为值类型时，可以直接用指针调用方法，反过来也同样可以
func (p person) test03() {
	p.name = "小邓"
	fmt.Println("test03()", p.name)
}

func (p *person) test04() { //地址拷贝
	(*p).name = "小邓"
	fmt.Println("test04()", (*p).name)
}

func main() {
	p := person{"小米"}
	test01(p)
	test02(&p)

	(&p).test03() //传入的是&p，但仍然是值拷贝
	fmt.Println("main() p.name=", p.name)

	(&p).test04()
	fmt.Println("main() p.name=", p.name) //形式上是传入值类型，但本质上是地址拷贝
	/*不管形式是怎么样，真正决定的是值拷贝还是地址拷贝由方法和那个类型绑定的 *pointer类型和value类型
	比如(p person)和(p *person)
	*/
}
