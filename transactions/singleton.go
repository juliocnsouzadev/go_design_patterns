package transactions

import (
	"bufio"
	"os"
	"strconv"
	"sync"
)

var (
	once     sync.Once
	instance Database
)

type Database interface {
	GetIncomesSum() float64
	GetExpensesSum() float64
}

type singletonDatabase struct {
	transactions []Transaction
}

func (db *singletonDatabase) GetIncomesSum() float64 {
	sum := 0.0
	for _, transaction := range db.transactions {
		if !transaction.In {
			continue
		}

		sum += transaction.Value
	}
	return sum
}

func (db *singletonDatabase) GetExpensesSum() float64 {
	sum := 0.0
	for _, transaction := range db.transactions {
		if transaction.In {
			continue
		}
		sum += transaction.Value * -1
	}
	return sum
}

func readData(path string) ([]Transaction, error) {
	ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(ex + path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := []Transaction{}

	for scanner.Scan() {
		var transaction Transaction
		transaction.Desc = scanner.Text()
		scanner.Scan()
		v, _ := strconv.ParseFloat(scanner.Text(), 64)
		transaction.Value = v
		scanner.Scan()
		tType := scanner.Text()
		if tType == "in" {
			transaction.In = true
		}
		result = append(result, transaction)
	}

	return result, nil
}

func GetSingletonDatabase() Database {
	once.Do(func() {
		db := singletonDatabase{}
		transactions, err := readData(".\\transactions.txt")
		if err == nil {
			db.transactions = transactions
		}
		instance = &db
	})
	return instance
}
