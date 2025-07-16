package main

import (
	"fmt"
)

// Get input from user
func getInput(prompt string) float64 {
	var input float64
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Invalid input, please enter a number.")
		return getInput(prompt)
	}
	return input
}

func main() {

	var revenue float64 = getInput("Enter the revenue amount: ")
	var expenses float64 = getInput("Enter the expenses amount: ")
	var taxRate float64 = getInput("Enter the tax rate (in %): ")

	profit := revenue - expenses
	tax := profit * (taxRate / 100)
	earningsBeforeTax := profit
	earningsAfterTax := earningsBeforeTax - tax
	ratio := earningsBeforeTax / profit

	fmt.Printf("Profit: $%.2f\n", profit)
	fmt.Printf("Tax: $%.2f\n", tax)
	fmt.Printf("Earnings Before Tax: $%.2f\n", earningsBeforeTax)
	fmt.Printf("Earnings After Tax: $%.2f\n", earningsAfterTax)
	fmt.Printf("Earnings Before Tax to Profit Ratio: %.2f%%\n", ratio*100)
}
