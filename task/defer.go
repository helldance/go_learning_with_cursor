package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Defer closing the file until the function finishes
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		} else {
			fmt.Println("File closed successfully")
		}
	}()

	// Read the file using bufio.Scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Print each line of the file
	}

	// Check for errors while reading the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
