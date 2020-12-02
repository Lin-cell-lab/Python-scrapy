package main

import (
	"fmt"
)

type usb interface {
	Start()
	End()
}

type phone struct {
}

func (phone phone) Start() {
	fmt.Println("手机开始工作。。。")
}

func (phone phone) End() {
	fmt.Println("手机结束工作。。。")
}

type camera struct {
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

func (c computer) Working(usb usb) {
	//通过接口变量调用start（）和end（）方法
	usb.Start()
	usb.End()
}

func main() {
	//创建结构体变量
	computer := computer{}
	phone := phone{}
	camera := camera{}

	//
	computer.Working(phone)
	computer.Working(camera)
}
