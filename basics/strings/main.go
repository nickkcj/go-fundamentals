package main

import "fmt"

// Creating a basic string 
var basicString string = "Hello, World!"

// Main function to print the basic string
func main() {
	fmt.Println("Basic String:", basicString)

	// Formatted output using Printf
	// %s is a placeholder for the string value, we can use it to format the output
	fmt.Printf("Formatted String: %s\n", basicString)

	// Using placeholder with .2f to format a float
	var number float64 = 3.14159
	fmt.Printf("Formatted Number: %.2f\n", number)

	// Using Sprintf to create a formatted string
	formattedString := fmt.Sprintf("Formatted String with Sprintf: %s", basicString)
	fmt.Println(formattedString)

	// If we want to use multiline strings, we can use backticks
	multilineString := `This is a multiline string.
	It can span multiple lines without needing to use \n.`
	fmt.Println("Multiline String:", multilineString)

	// We can also use a lot of placeholders in Printf
	fmt.Printf("Formatted String with multiple placeholders: %s, Length: %d\n", basicString, len(basicString))

	// Operations possible with strings
	fmt.Println("String Length:", len(basicString)) // Length of the string
	fmt.Println("First Character:", string(basicString[0])) // Accessing the first character
	fmt.Println("Substring (0-5):", basicString[0:5]) // Substring from index 0 to 5

	// Concatenation of strings
	concatenatedString := basicString + " How are you?"
	fmt.Println("Concatenated String:", concatenatedString)

	// Printing the type of the variable
	fmt.Printf("Type of basicString: %T\n", basicString)
}
