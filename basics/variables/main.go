// How to define variables and its types in Go
package main

import (
	"fmt"
)

// Defining a variable with a specific type
var myVariable int = 10

// Defining a variable without a specific type (type inference)
var myOtherVariable = "Hello, Go!"

const myConstant = 3.14

// Defining a variable using the short variable declaration syntax (only inside functions)
func main() {
	fmt.Println("myVariable:", myVariable)
	fmt.Println("myOtherVariable:", myOtherVariable)
	myShortVariable := 3.14
	fmt.Println(myShortVariable)
	fmt.Println("myConstant:", myConstant)
}

// This code demonstrates how to define variables and constants in Go, including type inference and short variable declaration syntax.