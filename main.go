package main

import (
	"fmt"
	"go_design_patterns/transactions"
)

func main() {
	fmt.Println("** Design Patterns in Go **\n")

	fmt.Println(">= Decorator Pattern <=\n************************")
	transaction := transactions.BankTransction
	transaction = transactions.SmsDecorator(transaction)
	transaction = transactions.EmailDecorator(transaction)
	transaction("salary", 1000, true)
	transaction("rent", 500, false)
	fmt.Println("************************\n")
}
