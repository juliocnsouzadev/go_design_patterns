package transactions

type TransctionFunc func(string, float64, bool ) float64

func BankTransction (desc string, value float64, in bool ) float64{
	if in {
		return value
	}	
	return value * -1
}
