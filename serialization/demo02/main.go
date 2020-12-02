package main

import (
	"encoding/json"
	"fmt"
)

//定义声明结构体
type monster struct {
	Name     string //自定义tag标签 序列化 monster小写不能跨包使用
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

//将json字符串 反序列化为struct
func unmarshStruct() {
	//说明str ，
	str := "{\"Name\":\"孙悟空\",\"Age\":21,\"Birthday\":\"2010-21-21\",\"Sal\":80000,\"Skill\":\"火眼金睛\"}"

	var monster monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err = %v\n", err)
	}
	fmt.Printf("反序列化后monster=%v\n", monster)
}

//json字符串反序列化为map
func unMarshMap() {
	str := "{\"address\":\"洪崖洞\",\"age\":123,\"name\":\"红孩儿\"}"

	//定义一个map
	var a map[string]interface{} //map反序列化无需make,make被封装到unmasrsh内部
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err = %v\n", err)
	}
	fmt.Printf("反序列化后[a]=%v\n", a)
}

func unMarshSlice() {
	str := "[{\"address\":\"介休\",\"age\":20,\"name\":\"小米\"},{\"address\":\"广安\",\"age\":21,\"name\":\"小邓\"}]"

	//定义一个切片
	var slice []map[string]interface{}
	//切片反序列化无需make,make被封装到unmasrsh内部
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err = %v\n", err)
	}
	fmt.Printf("反序列化后[a]=%v\n", slice)
}

//反序列化后的数据类型必须和序列化前一样
//如果json字符串是通过程序获取，则不需要转义处理
func main() {
	unmarshStruct()
	unMarshMap()
	unMarshSlice()
}
