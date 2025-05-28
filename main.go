package main

import (
	"fmt"
	"strconv"
)

// В цикле спросить о транзакциях, число -10, 10,40.5 добавлять в массив и вывести

func main() {
	transactions := []float64{}
	var result float64

	fmt.Println("Подсчет баланса: ")

	transactions, result = loopInputTransactions()

	outputResult(transactions, result)
}

func calculateResult(transactions []float64) float64 {
	var result float64
	for _, value := range transactions {
		result += value
	}
	return result
}

func breakeApp(word string) (breaked bool, trans float64) {
	num, err := strconv.ParseFloat(word, 64)
	if err == nil {
		return true, num
	}

	return false, 0.0
}

func loopInputTransactions() (transactions []float64, result float64) {
	var word string
	for {
		fmt.Println("Напиши число или любой символ для подведения итога")
		fmt.Scan(&word)
		breaked, transaction := breakeApp(word)
		if !breaked {
			break
		}
		transactions = append(transactions, transaction)
		result = calculateResult(transactions)
	}
	return transactions, result
}

func outputResult(transactions []float64, result float64) {
	fmt.Printf("Ваш баланс: %0.1f\n", result)
	fmt.Println(transactions)
}
