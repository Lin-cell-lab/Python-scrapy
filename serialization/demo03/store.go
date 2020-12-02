package main

import (
	"encoding/json"
	"fmt"
)

type monster struct {
	Name  string
	Age   int
	Skill string
}

//序列化
func (m *monster) Store() string {
	monster := monster{
		Name:  "六魔王",
		Age:   213,
		Skill: "打麻将",
	}
	date, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err = %v", err)
	}
	fmt.Printf("序列化成功，mosnter序列化后为%v", string(date))
	return string(date)
}

//反序列化
func (m *monster) ReStore() monster {
	str := "{\"Name\":\"六魔王\",\"Age\":213,\"Skill\":\"打麻将\"}"

	var monster monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("ReStore err = %v\n", err)
	}
	fmt.Printf("反序列化后monster=%v\n", monster)
	return monster
}
func main() {
	var a monster
	a.ReStore()
	a.Store()
}
