package utils

import "fmt"

// GetTwoNumbers prompts for and returns two numbers from user
func GetTwoNumbers() (float64, float64) {
	var num1, num2 float64

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	return num1, num2
}

// GetChoice gets the user's menu selection
func GetChoice() int {
	var choice int
	fmt.Print("Your choice: ")
	fmt.Scanln(&choice)
	return choice
}

// ShowResult displays the calculation result
func ShowResult(operation string, num1, num2, result float64) {
	fmt.Printf("\n%.2f %s %.2f = %.2f\n\n", num1, operation, num2, result)
}

// ShowError displays an error message
func ShowError(message string) {
	fmt.Printf("\nError: %s\n\n", message)
}
