package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test() hello world!" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test() //开启一个协程 (主线程执行完了，协程未执行完毕也要退出；协程执行完可自行退出！)
	for i := 1; i <= 10; i++ {
		fmt.Println("main() hello golang!" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
