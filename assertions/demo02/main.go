package main

import (
	"fmt"
)

type usb interface {
	Start()
	End()
}

type phone struct {
	Name string
}

func (phone phone) Start() {
	fmt.Println("手机开始工作。。。")
}

func (phone phone) End() {
	fmt.Println("手机结束工作。。。")
}

func (phone phone) Call() {
	fmt.Println("手机电话响起！")
}

type camera struct {
	Name string
}

func (camera camera) Start() {
	fmt.Println("相机开始工作。。。")
}

func (camera camera) End() {
	fmt.Println("相机结束工作。。。")
}

type computer struct {
}

//编写一个方法，接受一个usb接口变量
//所谓实现了usb接口，就是指实现了usb接口声明的所有方法
func (c computer) Work() {

}

func (c computer) Working(usb usb) {
	//通过接口变量调用start（）和end（）方法
	usb.Start()
	//如果usb指向的是phone结构体变量，则还需要调用call方法
	//e类型断言
	if phone, ok := usb.(phone); ok {
		phone.Call()
	}
	usb.End()
}

func main() {
	//定义一个usb接口数组，可以存放phone和camera变量
	var usbArr [3]usb
	usbArr[0] = phone{"小米"}
	usbArr[1] = phone{"华为"}
	usbArr[2] = camera{"尼康"}
	fmt.Println(usbArr)

	var computer computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}
}
