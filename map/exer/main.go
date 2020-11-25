package main

import (
	"fmt"
)

//信息录入系统
func modifyUser(users map[string]map[string]string, name string, nickname string) {
	//判断user中是否有name
	if users["name"] != nil {
		users["name"]["pwd"] = "8888888"
	} else {

		users[name] = make(map[string]string, 2)
		users[name]["nickname"] = nickname
		users[name]["pwd"] = "888888"
	}

}

func main() {

	users := make(map[string]map[string]string, 10)

	modifyUser(users, "小米", "米")
	fmt.Println(users)
}
