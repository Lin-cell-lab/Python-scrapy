package main

import (
	"encoding/json"
	"fmt"
)

//定义声明结构体
type monster struct {
	Name     string  `json:"monster_name"` //自定义tag标签 序列化 monster小写不能跨包使用
	Age      int     `json:"monster_age"`
	Birthday string  `json:"monster_birthday"`
	Sal      float64 `json:"monster_sal"`
	Skill    string  `json:"monster_skill"`
}

func test() {
	monster1 := monster{
		Name:     "孙悟空",
		Age:      21,
		Birthday: "2010-21-21",
		Sal:      80000,
		Skill:    "火眼金睛",
	}
	//返回一个字节列表
	date, err := json.Marshal(&monster1)
	if err != nil {
		fmt.Printf("序列号错误 err= %v", err)
	}
	fmt.Printf("[monster]序列化后=%v\n", string(date))
}

//map（映射）序列化
func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["address"] = "洪崖洞"
	a["name"] = "红孩儿"
	a["age"] = 123

	date, err := json.Marshal(&a)
	if err != nil {
		fmt.Printf("序列号错误 err= %v", err)
	}
	fmt.Printf("[aMap]序列化后=%v\n", string(date))
}

//切片序列化
func testSlice() {
	//结构体切片
	var slice []map[string]interface{}
	//结构体映射
	var m1 map[string]interface{}

	m1 = make(map[string]interface{})
	m1["name"] = "小米"
	m1["age"] = 20
	m1["address"] = "介休"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "小邓"
	m2["age"] = 21
	m2["address"] = "广安"
	slice = append(slice, m2)
	//将切片进行序列化操作
	date, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列号错误 err = %v", err)
	}

	fmt.Println("[slice]序列化后：\n", string(date))
}

//普通数据类型的序列化 ：
func testFloat64() {
	var num float64 = 12.21
	date, err := json.Marshal(num)
	if err != nil {
		fmt.Printf("序列号错误 err = %v", err)
	}
	fmt.Println("基本数据类型序列化为：", string(date))

}

func main() {
	test()
	testMap()
	testSlice()
	testFloat64()
}
