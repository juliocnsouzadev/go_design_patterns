package transactions

type Transaction struct {
	Desc  string
	Value float64
	In    bool
}

type TransactionFunc func(transaction Transaction) float64

func BankTransaction(transaction Transaction) float64 {
	if transaction.In {
		return transaction.Value
	}
	return transaction.Value * -1
}
