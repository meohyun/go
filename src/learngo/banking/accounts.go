package accounts

import "fmt"

// Acoount struct
type Account struct {
	owner   string // 대문자면 public
	balance int    // 소문자면 private
}

// NewAccount create
// account의 메모리주소를 return함.
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit amount on your account
// Method > struct의 private 변수에 접근할 수 있다.
func (a Account) Deposit(amount int) {
	a.balance += amount
	fmt.Println(a.balance)
}
