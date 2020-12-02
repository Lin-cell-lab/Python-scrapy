package main

import (
	"fmt"
	"runtime"
)

func main() {
	//查看cpu的数量
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum = ", cpuNum)

	//设置使用的cpu数量
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
