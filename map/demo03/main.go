package main

import (
	"fmt"
)

//map切片【map个数就动态化】
func main() {
	//1.声明一个map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)

	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "2222"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "孙悟空"
		monsters[1]["age"] = "5888"
	}

	//合理增加方式
	monster1 := map[string]string{
		"name": "白骨精",
		"age":  "200",
	}

	monsters = append(monsters, monster1)
	fmt.Println(monsters)
}
