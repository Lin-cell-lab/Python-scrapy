package main

import (
	"fmt"
)

//判断函数的输入类型
//空接口
func typeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数为bool类型，值为%v\n", index, x)
		case float32:
			fmt.Printf("第%v个参数为float32类型，值为%v\n", index, x)
		case float64:
			fmt.Printf("第%v个参数为float64类型，值为%v\n", index, x)
		case int, int32, int64, int8:
			fmt.Printf("第%v个参数为int类型，值为%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数为string类型，值为%v\n", index, x)
		case stu:
			fmt.Printf("第%v个参数为stu类型，值为%v\n", index, x)
		case *stu:
			fmt.Printf("第%v个参数为*stu类型，值为%v\n", index, x)
		default:
			fmt.Printf("第%v个参数类型不确定，值为%v\n", index, x)
		}
	}
}

type stu struct {
}

func main() {
	n1 := 1.1
	n2 := "虾米"
	var stu1 stu
	var stu2 *stu

	typeJudge(n1, n2, stu1, stu2)

}
