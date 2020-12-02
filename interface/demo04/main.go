package main

import (
	"fmt"
)

type monkey struct {
	Name string
}

type birdInterface interface {
	Flying()
	Swim()
}

type fishInterface interface {
	Swim()
}

func (this1 *monkey) climbing() {
	fmt.Println(this1.Name, "生来会爬树！")
}

//littlemonkey结构体
type littleMonkey struct {
	monkey //继承
}

func (this1 *littleMonkey) Flying() {
	fmt.Println(this1.Name, "会飞！")
}

func (this1 *littleMonkey) Swim() {
	fmt.Println("会游泳！")
}

func main() {
	//当A结构体继承了B结构体，那么A结构体就会自动继承B结构体的方法和字段，并且可以直接使用！
	//A需要拓展功能时，实现某个接口即可，接口是对继承的补充。
	//创建一个littlemonkey实例
	//继承：解决代码的复用性和可维护性
	//接口：设计好各种规范，让其他自定义类型去实现这些方法
	//接口比继承更加灵活 继承是 is - a ，接口只需满足like - a的关系 。接口在一定程度上解决代码的耦合性
	monkey := littleMonkey{
		monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swim()
}
