package main

import "fmt"

func main() {
	ch := make(chan int, 100) // Buffered channel with a capacity of 100

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(<-ch) // Output: 1
	fmt.Println(<-ch) // Output: 2
	fmt.Println(<-ch) // Output: 3
}
