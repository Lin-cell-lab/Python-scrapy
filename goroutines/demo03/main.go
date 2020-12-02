package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	//lock是一个全局的互斥锁
	//sync 是包，synchornzed同步
	lock sync.Mutex
)

func test(n int) {
	//计算n！，放入到myMap
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//将res放入myMap
	//加锁
	lock.Lock()
	myMap[n] = res //concurrent map writes
	//解锁
	lock.Unlock()

}

func main() {
	//计算1-200的各个数的阶乘，并且把各个数的阶乘放入map当中，使用goroutine完成

	//1.编写函数，计算各个数的阶乘，放入map当中 ，2. 启动的协程多个，放入map当中
	//错误： Found 3 data race(s)
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	//休眠十秒钟
	time.Sleep(time.Second * 5)

	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%v]:%v\n", i, v)
	}
	lock.Unlock()
}
