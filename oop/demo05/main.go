package main

import (
	"encoding/json"
	"fmt"
)

type a1 struct {
	num int
}

//
type b2 struct {
	num int
}

type monster struct {
	Name  string   `json:"name"` //struct tag标签 将大写名字换成小写
	Age   int      `json:"age"`
	Skill []string `json:"skill"`
}

func main() {
	var a a1
	var b b2
	a = a1(b) //支持强转换，但是结构体的字段要求完全一样（名字，个数，数据类型）
	fmt.Println(a, b)

	mosnter1 := monster{
		Name:  "孙悟空",
		Age:   600,
		Skill: []string{"火眼金睛", "筋斗云"}}

	//将monster序列化为json格式
	jsonMonster1, err := json.Marshal(mosnter1)
	if err != nil {
		fmt.Println("序列化错误")
	}
	fmt.Println("jsonMonster1=", string(jsonMonster1))
}
