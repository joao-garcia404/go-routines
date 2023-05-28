package main

import (
	"fmt"
)

// Thread 1
func main() {
	channel := make(chan string)

	// Thread 2
	go func() {
		channel <- "Channel content" // Fill the channel with content
	}()

	msg := <-channel // Channel is now empty and it can be filled now

	fmt.Println(msg)
}
