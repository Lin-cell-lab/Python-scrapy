package main

import (
	"fmt"
)

func main() {
	cities := make(map[string]string)
	cities["no1"] = "上海"
	cities["no2"] = "重庆"
	cities["no3"] = "北京"
	fmt.Printf("cities有%v个\n", len(cities))
	fmt.Println(cities)

	//演示map查找
	val, ok := cities["no1"]
	if ok {
		fmt.Printf("有no1的key，且值为%v\n", val)
	} else {
		fmt.Printf("没有no1的key\n")
	}

	//1.修改
	cities["no1"] = "华盛顿"
	fmt.Println(cities)

	//2.删除,如果值不存在，则不进行操作，也不报错
	delete(cities, "no1")
	delete(cities, "no4")
	fmt.Println(cities)
	//2.1一次性删除所,遍历删除
	cities = make(map[string]string) //直接make一个新空间

	//2.增加
	cities["no4"] = "纽约"
	fmt.Println(cities)
}
