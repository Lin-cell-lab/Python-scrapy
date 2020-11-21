package main

import (
	"fmt"
)

//演示golang中字符串的基本使用
func main() {
	//1.string的基本使用
	var address string = "北京长城 110 hello world1"
	fmt.Println(address)

	//2.不能修改string的内容,字符串不可变

	//3.字符串的两种表示形式，双引号和会识别的转义字符 反引号
	//（以字符串的原生形式输出，包括换行和特殊字符，可以实现防攻击和输出源代码）
	var str2 string = `abc\abe`
	fmt.Println(str2)

	var str3 string = `package main

	import (
		"fmt"
	)
	
	//演示golang中字符串的基本使用
	func main() {
		//1.string的基本使用
		var address string = "北京长城 110 hello world1"
		fmt.Println(address)
	
		//2.不能修改string的内容
	
		//3.字符串的两种表示形式，双引号和会识别的转义字符 反引号
		//（以字符串的原生形式输出，包括换行和特殊字符，可以实现防攻击和输出源代码）
		str2 = "abc\abe"
		fmt.Println(str2)
		`
	fmt.Println(str3)

	//4.字符串的拼接,多字符串拼接乐园进行换行相加，但加号必须留在最后面
	str := "hello" + "deng" + "!"
	str += "hahaha"
	fmt.Println(str)

	//5.基本数据类型的默认值
	var a int
	var b string
	var isc bool
	var c float32
	var d float64
	//%v按照变量的值输出
	fmt.Printf("a=%d,b=%v,isc=%v,c=%f,d=%f", a, b, isc, c, d)
}
