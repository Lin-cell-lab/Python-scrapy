package main

import (
	"fmt"
)

/*!.case后是一个表达式（常量，变量，一个有返回值的函数）
2.每个表达式的数据类型，必须和switch中的数值类型一样
3.case后可以带多个表达式，使用逗号间隔，*/
func main() {
	//switch使用案例
	var key byte
	fmt.Println("请输入一个字符(a,b,c,d):")
	fmt.Scanf("%c", &key)

	switch key {
	case 'a':
		fmt.Println("今天星期一")
	case 'b':
		fmt.Println("今天星期二")
	case 'c':
		fmt.Println("今天星期三")
	default:
		fmt.Println("输入有误...")
	}

	//使用细节
	var n1 int32 = 20
	var n2 int32 = 20
	switch n1 {
	case 20, n2: //case后可以有多个表达式子
		fmt.Println("ok1")
	default:
		fmt.Println("ok2")
	}

	//switch后可以不带表达式，类似于if-else 分支来使用、
	var age int = 20

	switch {
	case age == 10:
		fmt.Println("age == 10")
	case age == 20:
		fmt.Println("age == 20")
	default:
		fmt.Println("未匹配到")
	}

	var score int = 93

	switch {
	case score > 90:
		fmt.Println("成绩优异！")
	case score >= 70 && score <= 90:
		fmt.Println("成绩良好！")
	case score >= 60 && score < 70:
		fmt.Println("成绩及格！")
	default:
		fmt.Println("不及格！")
	}

	//switch穿透 fallthrought
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough //默认只能穿透一层
	case 20:
		fmt.Println("ok2")
	default:
		fmt.Println("未匹配到！")
	}

}
