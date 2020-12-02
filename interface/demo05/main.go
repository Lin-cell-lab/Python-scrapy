package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//1.hero结构体
type hero struct {
	Name string
	Age  int
}

//2.声明hero结构体切片类型
type heroSlice []hero

//3.实现接口,（获取长度）
func (hs heroSlice) Len() int {
	return len(hs)
}

//给定一个排序标准 ，年龄从小到大
func (hs heroSlice) Less(i, j int) bool {
	return hs[i].Age > hs[j].Age
}

//交换elemnts
func (hs heroSlice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}

func main() {
	var intSlice = []int{5, 2, 1, 93, 12, 43}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//对一个结构体切片进行排序
	//测试
	var heros heroSlice
	for i := 0; i < 10; i++ {
		hero := hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		//将hero放入切片
		heros = append(heros, hero)
	}
	//调用sort包使用
	sort.Sort(heros)
	fmt.Println(heros)
}
