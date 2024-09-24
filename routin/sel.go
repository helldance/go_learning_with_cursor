package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(10 * time.Second)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- "Message from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(6 * time.Second):
		fmt.Println("Timeout")
	}
}
