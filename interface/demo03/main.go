package main

import (
	"fmt"
)

type bInterface interface {
	test01()
}

type aInterface interface {
	test02()
}

type cInterface interface {
	bInterface
	aInterface
	test03()
}

//要实现从，则需要将a和b都实现
type students struct {
}

func (stu students) test01() {
	fmt.Println("test01()...print")
}

func (stu students) test02() {
	fmt.Println("test02()...print")
}

func (stu students) test03() {
	fmt.Println("test03()...print")
}

func main() {
	var stu students
	var a aInterface = stu
	a.test02()
	var c cInterface = stu
	c.test02()
	c.test03()
	c.test01()

}
