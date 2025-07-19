package main

// Here we are importing the necessary packages. That is how 
// we can use the functions defined in other files. 
// go.mod file is used to manage dependencies.
import (
	"example/packages.com/calculator"
	"example/packages.com/menu"
	"example/packages.com/utils"
	"github.com/pkg/errors" // We can also use third-party packages like this
	// We need to search for the package name in Go modules,
	// Then run the command `go get github.com/pkg/errors`
	// It will automatically update the go.mod and go.sum files
	// (Use 'go get' to install packages from the command line.)
)

func main() {
	// Use the menu package
	menu.Welcome()

	for {
		// Show menu from menu package
		menu.Show()

		// Get user choice from utils package
		choice := utils.GetChoice()

		// Exit if user chooses 5
		if choice == 5 {
			menu.Goodbye()
			break
		}

		// Get two numbers from utils package
		num1, num2 := utils.GetTwoNumbers()

		// Perform calculation based on choice using calculator package
		switch choice {
		case 1:
			result := calculator.Add(num1, num2)
			utils.ShowResult("+", num1, num2, result)

		case 2:
			result := calculator.Subtract(num1, num2)
			utils.ShowResult("-", num1, num2, result)

		case 3:
			result := calculator.Multiply(num1, num2)
			utils.ShowResult("×", num1, num2, result)

		case 4:
			result, err := calculator.Divide(num1, num2)
			if err != nil {
				utils.ShowError(err.Error())
			} else {
				utils.ShowResult("÷", num1, num2, result)
			}

		default:
			utils.ShowError("Invalid choice. Please choose 1-5.")
		}
	}
}
