package utils

import "fmt"

// GetInput prompts the user for input and returns a float64 value
func GetInput(prompt string) float64 {
	var input float64
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Invalid input, please enter a number.")
		return GetInput(prompt)
	}
	return input
}

// GetStringInput prompts the user for string input
func GetStringInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}
