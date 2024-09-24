package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello, Goroutine!"
	}()

	message := <-ch // Receive the value from channel
	fmt.Println(message)
}
