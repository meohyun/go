package accounts

import (
	"errors"
	"fmt"
)

var errMoney = errors.New("You can't withdraw.")

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
// 그냥 Account를 사용하면 account의 복사본을 만들게 된다. strutct 앞에 *을 붙여주면 복사본을 만들지 않고 balance를 증가시킬 수 있다.
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw amount in your acoount
// error는 error과 nil을 리턴할 수 있다. nil은 None에 해당한다.
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner your account
func (a *Account) ChangeOwner(changeowner string) {
	a.owner = changeowner
}

// Owner your account
func (a Account) Owner() string {
	return a.owner
}

// Go가 자동으로 호출해주는 메소드인 String 이다.
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}
