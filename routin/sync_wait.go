package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal the WaitGroup that this goroutine is done
	fmt.Printf("Worker %d starting\n", id)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait() // Wait for all workers to finish
	fmt.Println("All workers done")
}
