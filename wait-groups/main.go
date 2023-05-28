package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running \n", i, name)
		time.Sleep(200 * time.Millisecond)
		wg.Done()
	}
}

// Thread 1
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(20)

	//Thread 2
	go task("A", &waitGroup)

	// Thread 3
	go task("B", &waitGroup)

	waitGroup.Wait()
}
