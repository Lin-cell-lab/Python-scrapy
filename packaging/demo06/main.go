package model

import (
	"fmt"
)

//设置账户结构
type Account struct {
	AccountNo int
	balance   float64
	pwd       string
}

//获取账户
func (account *Account) SetAccount(account1 int) {
	if account1 > 6 && account1 < 10 {
		account.AccountNo = account1
	} else {
		fmt.Println("账号格式输入错误！")
	}
}

func (account *Account) SetBalanece(balance1 float64) {
	if balance1 > 20.0 {
		account.balance = balance1
	} else {
		fmt.Println("余额输入错误！")
	}
}

func (account *Account) SetPwd(pwd1 string) {
	if len(pwd1) == 6 {
		account.pwd = pwd1
	} else {
		fmt.Println("密码输入错误！")
	}
}

func (account *Account) GetAccountNo() int {
	return account.AccountNo
}

func (account *Account) GetBalanece() float64 {
	return account.balance
}

func (account *Account) GetPwd() string {
	return account.pwd
}
