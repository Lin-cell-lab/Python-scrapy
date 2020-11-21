package main

import "fmt"

func main() {
	var c3 int = '北' //整数型字符
	fmt.Printf("c3=%c,c3对应的码值%d", c3, c3)
	//var c3 byte = '北' //overflow益处

	var c1 byte = 'a'
	fmt.Println("c1=", c1) //直接输出byte值，相当于直接输出对应的字符码值

	var c4 int = 22269 //22269对应“国”
	fmt.Printf("%c", c4)

	var n1 = 10 + 'a' //字符类型可以运算，相当于整数 97+10=107
	fmt.Println("n1=", n1)
}

//保存的码值如果在ascii中，直接用byte表示
//如果保存的字符超出码值，大于255，用int保存
//格式化输出码值 fmt.println("%d",c3)
