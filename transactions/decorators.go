package transactions

import (
	"fmt"
)

func SmsDecorator(transaction TransactionFunc) TransactionFunc {
	return func(trans Transaction) float64 {
		fn := func(trans Transaction) float64 {
			tType := "Out"
			if trans.In {
				tType = "In"
			}
			defer func() {
				fmt.Printf("SMS Sent\tTransaction [Value: %v, type: %s] - %s\n", trans.Value, tType, trans.Desc)
			}()
			return transaction(trans)
		}
		return fn(trans)
	}
}

func EmailDecorator(transaction TransactionFunc) TransactionFunc {
	return func(trans Transaction) float64 {
		fn := func(trans Transaction) float64 {
			tType := "Out"
			if trans.In {
				tType = "In"
			}
			defer func() {
				fmt.Printf("Email Sent\tTransaction [Value: %v, type: %s] - %s\n", trans.Value, tType, trans.Desc)
			}()
			return transaction(trans)
		}
		return fn(trans)
	}
}
