package main 

// It is common in Go that functions return an error as the last return value
// This allows the caller to check if the operation was successful or if an error occurred
import (
	"fmt"
	"os"
	"errors"
)

func main() {
	content, err := readFile("nonexistent.txt")
	if err != nil {
		fmt.Println(err)
		panic("Can't continue. Sorry :(") // OR return... If an error occurs, we print it and exit the program
		// In a real application, you might want to handle the error differently,
		// such as logging it or retrying the operation
		// Here, we simply print the error message and exit
	} else {
		fmt.Println("File content:", content)
	}
}

// Function to read a file and handle errors
func readFile(filename string) (string, error) {
	// os.ReadFile reads the content of the specified file
	// It returns the content as a byte slice and an error if any
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", errors.New("failed to read file") // Wrap the error 
	}
	return string(content), nil // Convert the byte slice to a string and return it
}