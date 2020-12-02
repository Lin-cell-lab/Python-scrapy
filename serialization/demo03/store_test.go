package main

import (
	"testing"
)

func TestStore(t *testing.T) {
	var monster1 monster

	str := "{\"Name\"：\"六魔王\",\"Age\":213,\"Skill\":\"打麻将\"}"
	m1 := monster1.Store()
	if m1 == str {
		t.Fatalf("Store()执行错误，期望序列为%v,实际序列为%v", str, m1)
	}
	t.Logf("Store()执行正确！")

}

func TestReStore(t *testing.T) {
	var monster2 monster
	//
	str := &monster{
		Name:  "牛魔王",
		Age:   213,
		Skill: "打麻将",
	}
	m2 := monster2.ReStore()
	if m2 == *str {
		t.Fatalf("ReStore()执行错误，期望序列为%v,实际序列为%v", str, m2)
	}
	t.Logf("ReStore()执行正确！")
}
