package model

import (
	"fmt"
)

type person struct {
	Name   string
	age    int
	salary float64
}

//构造函数
func newPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//get获取
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄输入不正确")
		p.age = 0
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) Setsalary(sal float64) {
	if sal > 0 && sal < 10000 {
		p.salary = sal
	} else {
		fmt.Println("薪水输入不正确")
		p.salary = 0
	}

}

//返回
func (p *person) GetSalary() float64 {
	return p.salary
}
