package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//如何生成0-100的整数(生成一个随机数，还需要以恶搞rand种子)

	fmt.Println(time.Now().UnixNano()) //种子设置，.UnixNano为纳秒/Unix为妙

	var count int = 0
	rand.Seed(time.Now().UnixNano()) //将种子放在循环外面

	for {
		n := rand.Intn(10000) + 1
		count++
		fmt.Println(n)

		if n == 99 {
			fmt.Printf("一共花了%v生成了99", count)
			break
		}
	}
}
