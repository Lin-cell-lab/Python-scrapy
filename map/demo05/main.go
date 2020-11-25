
package main

import (
	"fmt"
)

func modify(map1 map[int]int) {
	map1[25] = 2121
}

func main() {
	//1.map为引用类型，遵守引用类型传递机制，一个函数接受map时，修改后，会直接修改原来的map
	map1 := make(map[int]int, 10)
	map1[25] = 44
	map1[11] = 52
	map1[42] = 71
	map1[32] = 51
	map1[10] = 53
	fmt.Println(map1)
	modify(map1)
	fmt.Println(map1)

	//2.map的容量达到后，在想增加map元素，会自动扩容，不会发生panic，map自动增长

	//3.map的value经常使用struct类型，管理结果更加复杂的数据（一个对象）
}