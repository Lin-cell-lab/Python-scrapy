package main

import (
	"fmt"
)

//1定义一个银行账户
type account struct {
	accountNo string
	Pwd       string
	Balance   float64
}

//2.存款
func (account *account) Deposite(money float64, pwd string) {
	//验证密码
	if pwd != account.Pwd {
		fmt.Println("输入密码不对!")
		return
	}
	//验证存款
	if money <= 0 {
		fmt.Println("输入金额不对!")
		return
	}
	account.Balance += money
	fmt.Println("存款成功！")
}

//3.取款
func (account *account) withDraw(money float64, pwd string) {
	//验证密码
	if pwd != account.Pwd {
		fmt.Println("输入密码不对!")
		return
	}
	//验证存款
	if money <= 0 || money > account.Balance {
		fmt.Println("输入金额不对!")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功！")
}

//4.查询
func (account *account) query(pwd string) {
	//验证密码
	if pwd != account.Pwd {
		fmt.Println("输入密码不对!")
		return
	}
	//查看金额
	fmt.Printf("您的账号为[%v],您的余额还有[%v]", account.accountNo, account.Balance)
}
func main() {
	account1 := account{
		accountNo: "213",
		Pwd:       "21213",
		Balance:   213.1,
	}

	var option int
	var pwd1 string
	var money1 float64
	for {
		fmt.Println("请选择你需要的操作：【1】：存款  【2】：取款  【3】：查询")
		fmt.Scanln(&option)
		switch option {
		case 1:
			fmt.Println("您选择的是【1】：存款")
			fmt.Println("请输入密码：")
			fmt.Scanln(&pwd1)
			fmt.Println("请输入金额：")
			fmt.Scanln(&money1)
			account1.Deposite(money1, pwd1)
		case 2:
			fmt.Println("您选择的是【2】：取款")
			fmt.Println("请输入密码：")
			fmt.Scanln(&pwd1)
			fmt.Println("请输入金额：")
			fmt.Scanln(&money1)
			account1.withDraw(money1, pwd1)
		case 3:
			fmt.Println("您选择的是【3】：查询")
			fmt.Println("请输入密码：")
			fmt.Scanln(&pwd1)
			account1.query(pwd1)
		default:
			fmt.Println("输入有误，请重新输入！")
		}
	}

}
