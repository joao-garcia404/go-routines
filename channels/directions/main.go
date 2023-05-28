package main

import "fmt"

// Send only channel type
func receive(name string, content chan<- string) {
	content <- name
}

// Receive only channel type
func read(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	content := make(chan string)

	go receive("info", content)
	read(content)
}
