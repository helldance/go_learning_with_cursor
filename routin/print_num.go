package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go printNumbers() // Runs in a separate goroutine
	fmt.Println("Goroutine started")

	time.Sleep(1 * time.Second) // Give time for goroutine to finish
}
