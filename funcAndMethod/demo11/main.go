package main

import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {

	return func(name string) string {
		//如果无指定的后缀，则加上，否则返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}

func main() {
	f := makeSuffix(".jpg")
	fmt.Println("文件处理后为", f("winter"))
}
