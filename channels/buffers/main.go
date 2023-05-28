package main

func main() {
	ch := make(chan string, 2) // Channel can now receive two messages without the first one been read

	ch <- "Content"
	ch <- "Content 2"

	println(<-ch)
	println(<-ch)
}
