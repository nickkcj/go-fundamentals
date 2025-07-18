package main

import (
	"fmt"
	"os"
)

// Results struct to organize our calculations
type Results struct {
	Profit            float64
	Tax               float64
	EarningsBeforeTax float64
	EarningsAfterTax  float64
	Ratio             float64
}

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

// writeFile with error handling
func writeFile(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v\n", filename, err)
		return err
	}
	fmt.Printf("Successfully wrote to %s\n", filename)
	return nil
}

// Save all results to files
func saveResults(results Results) {
	writeFile("profit.txt", fmt.Sprintf("Profit: $%.2f\n", results.Profit))
	writeFile("tax.txt", fmt.Sprintf("Tax: $%.2f\n", results.Tax))
	writeFile("earnings_before_tax.txt", fmt.Sprintf("Earnings Before Tax: $%.2f\n", results.EarningsBeforeTax))
	writeFile("earnings_after_tax.txt", fmt.Sprintf("Earnings After Tax: $%.2f\n", results.EarningsAfterTax))
	writeFile("ratio.txt", fmt.Sprintf("Earnings Before Tax to Profit Ratio: %.2f%%\n", results.Ratio*100))
}

// Display results to console
func displayResults(results Results) {
	fmt.Printf("\n=== PROFIT CALCULATION RESULTS ===\n")
	fmt.Printf("Profit: $%.2f\n", results.Profit)
	fmt.Printf("Tax: $%.2f\n", results.Tax)
	fmt.Printf("Earnings Before Tax: $%.2f\n", results.EarningsBeforeTax)
	fmt.Printf("Earnings After Tax: $%.2f\n", results.EarningsAfterTax)
	fmt.Printf("Earnings Before Tax to Profit Ratio: %.2f%%\n", results.Ratio*100)
}

func main() {

	var revenue float64 = getInput("Enter the revenue amount: ")
	if revenue <= 0 {
		fmt.Println("Revenue must be a positive number.")
		return
	}
	var expenses float64 = getInput("Enter the expenses amount: ")
	if expenses <= 0 {
		fmt.Println("Expenses must be a positive number.")
		return
	}
	if expenses > revenue {
		fmt.Println("Expenses cannot exceed revenue.")
		return
	}
	var taxRate float64 = getInput("Enter the tax rate (in %): ")
	if taxRate < 0 || taxRate > 100 {
		fmt.Println("Tax rate must be between 0 and 100.")
		return
	}

	// Calculate results
	profit := revenue - expenses
	tax := profit * (taxRate / 100)
	earningsBeforeTax := profit
	earningsAfterTax := earningsBeforeTax - tax
	ratio := earningsBeforeTax / profit

	results := Results{
		Profit:            profit,
		Tax:               tax,
		EarningsBeforeTax: earningsBeforeTax,
		EarningsAfterTax:  earningsAfterTax,
		Ratio:             ratio,
	}

	// Save to files
	saveResults(results)

	// Display to console
	displayResults(results)
}
