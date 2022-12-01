package banking

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't Withdraw")

// 새로운 계좌를 만드는 Function
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// method이다! func (struct) 이런 식으로 쓴다.
// struct의 메써드로 쓰면 일단 디폴트로 복사본을 가져온다. 그렇기 때문에 그냥 바꾸면 안된다 이거야. 포인터를 가져와야한다.
func (theAccount *Account) Deposit(amount int) {
	theAccount.balance += amount
}

// getBalance method
func (a Account) Balance() int {
	return a.balance
}

//  Withdraw x amount from your account
// error를 리턴하는 법
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

func (a Account) ToString() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
