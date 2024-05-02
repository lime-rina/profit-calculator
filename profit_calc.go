package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate float64

	fmt.Print("Input your revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Input your expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Input your tax rate: ")
	fmt.Scan(&taxRate)

	EBT := revenue - expenses - (revenue-expenses)*taxRate/100
	profit := EBT - (EBT * taxRate / 100)
	ratio := EBT / profit

	fmt.Println("EBT: ", EBT)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
}
