package main

import (
	"fmt"
)

//Cat结构体，将各字段/属性信息，放入到cat结构体中进行管理（类）
type cat struct {
	Name  string
	Age   int
	Color string
	Hoby  string
}

func main() {

	//1.使用变量解决养猫问题
	var cat1 cat //对象
	fmt.Printf("cat1的地址为%p\n", &cat1)
	cat1.Name = "小米"
	cat1.Age = 20
	cat1.Color = "蓝色"
	cat1.Hoby = "吃鱼"
	fmt.Printf("cat.name的地址为%p\n", &cat1.Name)
	fmt.Printf("cat.age的地址为%p\n", &cat1.Age)
	fmt.Printf("cat.coloe的地址为%p\n", &cat1.Color)
	fmt.Println(cat1)

}
