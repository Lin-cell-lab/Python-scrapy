package main

import (
	"fmt"
)

//map映射基础
func main() {
	//1.声明方式，map定义时不分配数据空间
	var a map[string]string

	//2.使用make给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "李逵"
	a["no2"] = "宋江"
	a["no3"] = "武松"
	fmt.Println(a)

	//3.
	cities := make(map[string]string)
	cities["no1"] = "上海"
	cities["no2"] = "重庆"
	cities["no3"] = "北京"
	fmt.Println(cities)

	//4.
	heroes := map[string]string{
		"hero1": "超人",
		"hero2": "蝙蝠侠",
		"hero3": "钢铁侠",
	}
	fmt.Println(heroes)

	//5.案例
	studentInfo := make(map[string]map[string]string)
	studentInfo["stdu01"] = make(map[string]string, 3)
	studentInfo["stdu01"]["name"] = "小米"
	studentInfo["stdu01"]["gender"] = "女"
	studentInfo["stdu01"]["address"] = "介休"

	studentInfo["stdu02"] = make(map[string]string, 3)
	studentInfo["stdu02"]["name"] = "小邓"
	studentInfo["stdu02"]["gender"] = "男"
	studentInfo["stdu02"]["address"] = "邻水"

	fmt.Println(studentInfo)
	fmt.Println(studentInfo["stdu02"])

}
