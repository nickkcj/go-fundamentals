package menu

import "fmt"

// IMPORTANT: Functions that aren't in the same package, need to have their 
// first letter capitalized to be exported. This is how we can use them in other packages.
// If its going to be used only within this package, it can start with a lowercase letter.

// Show displays the main menu options
func Show() {
	fmt.Println("=== CALCULATOR WITH PACKAGES ===")
	fmt.Println("1. Add two numbers")
	fmt.Println("2. Subtract two numbers")
	fmt.Println("3. Multiply two numbers")
	fmt.Println("4. Divide two numbers")
	fmt.Println("5. Exit")
}

// Welcome displays a welcome message
func Welcome() {
	fmt.Println("Welcome to the Calculator!")
	fmt.Println("This example shows how to use DIFFERENT packages.")
	fmt.Println()
}

// Goodbye displays a goodbye message
func Goodbye() {
	fmt.Println("Thanks for using the calculator!")
	fmt.Println("Goodbye! 👋")
}
