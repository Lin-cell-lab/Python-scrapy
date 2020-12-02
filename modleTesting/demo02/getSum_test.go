package main

import "testing"

func TestGetSub(t *testing.T) {
	//测试模板
	res := getSub(10, 21)
	if res != 210 {
		t.Fatalf("getSub(10,21)执行错误，期望值=%v，实际值 = %v\n", 210, res)
	}
	t.Logf("getSub(10,21)，执行正确")
}
