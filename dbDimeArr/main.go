package main

import "fmt"

func main() {
	var arrStr [2][2]string
	arrStr[0][0] = "小米"
	arrStr[0][1] = "小邓"
	arrStr[1][0] = "小温"
	arrStr[1][1] = "小张"

	fmt.Println(arrStr)
}
