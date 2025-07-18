package main

import "os"

func main() {
	// writeFile("output.txt", "Hello, World!\nThis is a file created by Go.")
	content := readFile("output.txt")
	os.Stdout.WriteString(content) // Print the content to standard output
}

// We import "os" to handle file operations
// func writeFile(filename string, content string) {
// 	// os.WriteFile writes the content to the specified file
// 	// If the file does not exist, it will be created
// 	os.WriteFile(filename, []byte(content), 0644) // 0644 sets the file permissions, owner can read/write, others can read

// }

func readFile(filename string) string {
	// os.ReadFile reads the content of the specified file
	// It returns the content as a byte slice and an error if any
	// We can convert the byte slice to a string
	content, err := os.ReadFile(filename)
	if err != nil {
		return "Error reading file: " + err.Error()
	}
	// If the data is a float, we can use strconv.ParseFloat to convert it
	// If the data is a string, we can convert it directly to a string
	return string(content)

}
