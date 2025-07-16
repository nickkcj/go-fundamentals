package main

import (
	"fmt"
	"math"
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
	/// Here is some ways to declare variables in Go

	var investmentAmount float64 = getInput("Enter the amount you want to invest: ")
	// Or we could simply do fmt.Scan(&investmentAmount), pointer to the variable.

	// Declare a variable using the const keyword
	const inflationRate = 5.0

	years := getInput("Enter the number of years you plan to invest: ")
	expectedReturn := getInput("Enter the expected annual return rate (in %): ")

	futureValue := math.Round(investmentAmount * math.Pow((1+expectedReturn/100), years))
	futureRealValue := math.Round(futureValue / math.Pow((1+inflationRate/100), years)) // Assuming 2% inflation rate
	fmt.Println("Future Value of Investment: $", futureValue)
	fmt.Println("Future Real Value of Investment: $", futureRealValue)

}
