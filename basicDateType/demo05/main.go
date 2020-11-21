package main

import (
	"fmt"
	"strconv"
)

//string转基本数据类型
func main() {
	var str string = "true"
	var b bool
	//b, _ = strconv.ParseBool(str)
	//1.strconv.ParseBool(str)会返回两个值，bool值和err
	//2.只想获取bool，使用下划线忽略err
	b, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Print("转换错误")
	}
	fmt.Printf("b type is %T b=%v\n", b, b)

	//字符如hello不能转为整数，默认转为值为0
	str2 := "123456"
	var n1 int64
	n1, err1 := strconv.ParseInt(str2, 10, 64)
	if err1 != nil {
		fmt.Print("转换错误")
	}
	fmt.Printf("n1 type is %T,n1=%v\n", n1, n1)

	str3 := "1231.21"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type is %T f1=%v", f1, f1)

}
