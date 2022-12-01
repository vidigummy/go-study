package main

import (
	"fmt"
	"log"

	"example.com/m/v2/banking"
)

func main() {
	newAccount := banking.NewAccount("vidigummy")
	newAccount.Deposit(3000000)
	err := newAccount.Withdraw(100)
	if err != nil {
		log.Fatalln(err)
	}
	newAccount.ChangeOwner("류동인")
	fmt.Println(newAccount.ToString())
}
