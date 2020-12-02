package main

//1.channel是线程安全（多个goroutine访问时，不需要加锁，channel本质上就是线程安全）
//2.多个协程操作同一个管道时，不会发生资源竞争的问题。
//3.管道本质是一个数据结构，队列 数据采用先进先出法
//4.channel有类型，string只能存放string数据类型
func main() {
	// var intChan chan int
	// intChan = make(chan int , 3)
}
