package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//1.创建一个数组放置a-z，使用for循环打印出来
	var byt [26]byte

	//1.放置元素
	for i := 0; i < 26; i++ {
		byt[i] = 'A' + byte(i)
	}

	// for _, value := range byt {
	// 	fmt.Println(value)
	// }

	for i := 0; i < 26; i++ {
		fmt.Printf("%c\t\n", byt[i])
	}

	//2.求出最大值，获取对应下标
	//wojiaodengjianglin woshisb
	var intArr [5]int = [...]int{1, 5, 3, 8, 5}
	maxVal := intArr[0]
	maxValIndex := 0

	//i = 1 和另外四个比
	for i := 1; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("maxVal=%v,maxValIndex=%v", maxVal, maxValIndex)

	//3.求出一个数组的平均值与和
	var numArr [5]int = [...]int{1, 2, 3, 1, 76}
	sum := 0

	for _, value := range numArr {
		sum += value
	}

	fmt.Printf("numArr数组的和为%v，平均值为%v\n", sum, float64(sum)/float64(len(numArr)))

	//4.随机生成五个数，并将其反转打印
	var intArr2 [5]int
	//为了没出生成的随机数不一样，给一个seed值  :rand.Seed(time.Now().Unix())
	len := len(intArr2)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		intArr2[i] = rand.Intn(100) // ==0-<100
	}
	fmt.Println(intArr2)

	temp := 0
	for i := 0; i < len/2; i++ {
		temp = intArr2[len-1-i]
		intArr2[len-1-i] = intArr2[i]
		intArr2[i] = temp
	}
	fmt.Println(intArr2)
}
