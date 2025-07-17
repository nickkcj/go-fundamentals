package main

import "fmt"

var accountBalance float64 = 1000.00

func main() {
	for {
		fmt.Println("Welcome to the Bank Application!")
		fmt.Print(`What do you want to do?
1. Check Balance
2. Deposit Money
3. Withdraw Money
4. Exit
`)

		choice := getInput("Choice: ")
		if choice != 4 {
			bankApp(int(choice))
		} else if choice == 4 {
			fmt.Println("Thank you for using the Bank Application. Goodbye!")
			break
		}
	}
}

func bankApp(choice int) {
	switch choice {
	case 1:
		fmt.Printf("Your current balance is: $%.2f\n", accountBalance)
	case 2:
		deposit := getInput("Enter amount to deposit: ")
		if deposit > 0 {
			accountBalance += deposit
			fmt.Printf("You have successfully deposited $%.2f. New balance: $%.2f\n", deposit, accountBalance)
		} else {
			fmt.Println("Deposit amount must be positive.")
		}

	case 3:
		withdraw := getInput("Enter amount to withdraw: ")
		if withdraw > 0 && withdraw <= accountBalance {
			accountBalance -= withdraw
			fmt.Printf("You have successfully withdrawn $%.2f. New balance: $%.2f\n", withdraw, accountBalance)
		}

	case 4:
		fmt.Println("Thank you for using the Bank Application. Goodbye!")
	}
}
	

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