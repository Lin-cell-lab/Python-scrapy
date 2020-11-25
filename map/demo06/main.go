package main

import (
	"fmt"
)

//map遍历
func main() {
	cities := make(map[string]string)
	cities["no1"] = "上海"
	cities["no2"] = "重庆"
	cities["no3"] = "北京"
	fmt.Println(cities)
	//1.单层遍历
	for key, value := range cities {
		fmt.Printf("key=%v,value=%v\n", key, value)
	}

	//2.多层遍历
	studentInfo := make(map[string]map[string]string)
	studentInfo["stdu01"] = make(map[string]string, 3)
	studentInfo["stdu01"]["name"] = "小米"
	studentInfo["stdu01"]["gender"] = "女"
	studentInfo["stdu01"]["address"] = "介休"

	studentInfo["stdu02"] = make(map[string]string, 3)
	studentInfo["stdu02"]["name"] = "小邓"
	studentInfo["stdu02"]["gender"] = "男"
	studentInfo["stdu02"]["address"] = "邻水"

	for k1, v1 := range studentInfo {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\ti=%v,j=%v\n", k2, v2)
		}
		fmt.Println()
	}

}
