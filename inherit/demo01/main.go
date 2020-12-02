package main

import "fmt"

//共有属性
type student struct {
	Name  string
	Age   int
	Score float64
}

//共有方法(年龄和成绩)
func (stu *student) ShowInfo() {
	fmt.Printf("学生姓名：%v, 学生年龄：%v,学生成绩：%v", stu.Name, stu.Age, stu.Score)
}

func (stu *student) SetScore(score float64) {
	stu.Score = score
}

func (stu *student) testing() {
	fmt.Println("正在考试！")
}

//小学生
type pupil struct {
	student        //结构体嵌入匿名结构体
	class   string //新增属性
}

//小米特有的方法 微笑
func (sstu *pupil) smile() {
	fmt.Println("====小米每天要微笑哦！====")
}

func main() {

	pupil := &pupil{class: "高三十七班"}
	pupil.student.Name = "小米"
	pupil.student.Age = 20
	fmt.Printf("小米的班级为:%v", pupil.class)
	pupil.testing()
	pupil.student.SetScore(2.2)
	pupil.student.ShowInfo()
	pupil.smile()
}
