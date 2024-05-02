package main

import (
	"fmt"
)

const revenueMsg, expensesMsg, taxMsg = "Input your revenue: ",
	"Input your expenses: ",
	"Input your tax rate: "

func main() {
	var revenue, expenses, taxRate float64

	handleUserInput(revenueMsg, &revenue)
	handleUserInput(expensesMsg, &expenses)
	handleUserInput(taxMsg, &taxRate)

	EBT, profit, ratio := calculateAll(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", EBT)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
}

func handleUserInput(message string, value *float64) {
	fmt.Print(message)
	fmt.Scan(value)
}

func calculateAll(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses - (revenue-expenses)*taxRate/100
	profit = ebt - (ebt * taxRate / 100)
	ratio = ebt / profit

	return ebt, profit, ratio
}
