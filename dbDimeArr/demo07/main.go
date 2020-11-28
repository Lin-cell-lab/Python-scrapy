package main

import (
	"fmt"
)

//1.冒泡排序法进行排序
func bubbleSort(arr *[10]int) {
	fmt.Println("排序前arr=", (*arr))

	temp := 0
	for j := 0; j < len(*arr)-1; j++ {

		for i := 0; i < len(*arr)-1-j; i++ {
			if (*arr)[i] > (*arr)[i+1] {
				temp = (*arr)[i]
				(*arr)[i] = (*arr)[i+1]
				(*arr)[i+1] = temp
			} else {
				continue
			}
		}
	}
	fmt.Println("排序后arr=", (*arr))
}

//2.快速排序
func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	//定义左右指针
	left, right := 0, len(arr)-1
	//定义基准值
	flag := (arr)[0]

	for i := 1; i <= right; {
		if (arr)[i] > flag {
			(arr)[i], (arr)[right] = (arr)[right], (arr)[i]
			right--
		} else {
			(arr)[i], (arr)[left] = (arr)[left], (arr)[i]
			i++
			left++
		}
	}
	quickSort(arr[:left])
	quickSort(arr[left+1:])

}

//3.插入排序
/*func insertSort(arr [10]int, insertNum int) {

	for i := 1; i <= len(arr); i++ {
		temp := arr[i]
		for j := 1; j <= len(arr)-i; j++ {
			if j
		}
	}
}*/

//4.二分查找
func binaryFind(arr *[10]int, leftIndex int, rightIndex int, findValue int) {

	//判断leftIndex > rightIndex 找不到
	if leftIndex > rightIndex {
		fmt.Println("未找到")
		return
	}

	//找中间数的下标
	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findValue {
		//查找数在leftIndex - middle
		binaryFind(arr, leftIndex, middle-1, findValue)
	} else if (*arr)[middle] < findValue {
		//查找数在 middle + 1 - rightIndex
		binaryFind(arr, middle+1, rightIndex, findValue)
	} else {
		fmt.Printf("扎到了，下标为%v\n", middle)
	}

}

func main() {
	/*	var arr [10]int
		//1.生成随机数组
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < len(arr); i++ {
			n := rand.Intn(100) + 1
			arr[i] = n
		}

		//2.冒泡排序法进行排序
		bubbleSort(&arr)
		binaryFind(&arr, 0, len(arr)-1, 90)*/
	arr := make([]int, 10)
	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < 10; i++ {
	// 	arr[i] = rand.Intn(10)
	// }
	arr[0] = 7
	arr[1] = 0
	arr[2] = 1
	arr[3] = 8
	arr[4] = 6
	arr[5] = 2
	arr[6] = 4
	arr[7] = 9
	arr[8] = 2
	arr[9] = 4
	quickSort(arr)
	fmt.Println(arr)
}
