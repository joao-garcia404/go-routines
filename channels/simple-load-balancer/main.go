package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	data := make(chan int)
	workersAmount := 100000

	// Init workers
	for i := 1; i <= workersAmount; i++ {
		go worker(i, data)
	}

	for i := 0; i < 1000000; i++ {
		data <- i
	}
}
