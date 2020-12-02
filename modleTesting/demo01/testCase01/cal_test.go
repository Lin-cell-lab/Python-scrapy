package main

import (
	"fmt"
	"testing"
)

//编辑测试用例，测试upper
func TestUpper(t *testing.T) {
	//调用
	res := addUpper(3)
	if res != 6 {
		t.Fatalf("AddUpper(10)执行错误，期望值=%v，实际值 = %v\n", 6, res)
	}
	//运行正确，输出日志
	t.Logf("AddUpper(6)，执行正确")
}
 
func testHello(t *testing.T) {
	fmt.Println("testHello被调用。。。")

}
