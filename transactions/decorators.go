package transactions

import (
	"fmt"
)

func SmsDecorator(transaction TransctionFunc) TransctionFunc {
	return func(desc string, value float64, in bool) float64 {
		fn := func(desc string, value float64, in bool) float64 {
			tType := "Out"
			if in {
				tType = "In"
			}
			defer func() {
				fmt.Printf("SMS Sent\tTransaction [value: %v, type: %s] - %s\n", value, tType, desc)
			}()
			return transaction(desc, value, in)
		}
		return fn(desc, value, in)
	}
}

func EmailDecorator(transaction TransctionFunc) TransctionFunc {
	return func(desc string, value float64, in bool) float64 {
		fn := func(desc string, value float64, in bool) float64 {
			tType := "Out"
			if in {
				tType = "In"
			}
			defer func() {
				fmt.Printf("Email Sent\tTransaction [value: %v, type: %s] - %s\n", value, tType, desc)
			}()
			return transaction(desc, value, in)
		}
		return fn(desc, value, in)
	}
}
