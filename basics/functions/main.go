package main

import "fmt"

// Basic function
func greet(name string) string {
	return "Hello, " + name + "!"
}

// Function with multiple parameters
func add(a, b int) int {
	return a + b
}

// Function with multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Named return values
func rectangle(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // naked return
}

func main() {
	// Basic function call
	message := greet("Gopher")
	fmt.Println(message)

	// Multiple parameters
	sum := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", sum)

	// Multiple return values
	result, err := divide(10.0, 3.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10.0 / 3.0 = %.2f\n", result)
	}

	// Named return values
	area, perimeter := rectangle(5.0, 3.0)
	fmt.Printf("Rectangle: area=%.2f, perimeter=%.2f\n", area, perimeter)
}
