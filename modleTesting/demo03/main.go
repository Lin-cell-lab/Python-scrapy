package main

import "fmt"

//测试函数
func addUpper(n int) int {
	res := 0
	for i := 1; i < n; i++ {
		res += i
	}
	return res
}

func main() {
	//测试方法一
	//缺点：需要在main函数中调用，肯能需要停止项目进行测试
	//不利于管理，混杂逻辑 ---
	res := addUpper(10) //55 == 55 函数正确
	fmt.Println("res=", res)
	//方法二，测试框架
	//1.单个测试文件，只需加上文件名 ，go test -v cal_test.go
	//2.测试单个方法，go test -v -test.run TestAddUpper
}

