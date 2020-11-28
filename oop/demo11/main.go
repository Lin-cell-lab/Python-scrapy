package main

import (
	"fmt"
)

type students struct {
	name   string
	gender string
	age    int
	id     int
	score  int
}

func (stu *students) say() string {
	info := fmt.Sprintf("students的信息：name[%v]\ngender[%v]\nage[%v]\nid[%v]score[%v]",
		stu.name,
		stu.gender,
		stu.age,
		stu.id,
		stu.score)
	return info
}

func main() {
	var stu1 students
	stu1.name = "小米"
	stu1.age = 12
	stu1.gender = "女"
	stu1.id = 1231312
	stu1.score = 21
	say1 := stu1.say()
	fmt.Println(say1)

	//返回结构体类型指针（（stu--> , 地址-->）
	var stu6 = &students{
		name: "小米",
	}
	// dengjianglinshigediaomao xianrennueduozhe
	fmt.Println(stu6.name)
}
