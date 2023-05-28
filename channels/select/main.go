package main

import "time"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 2
	}()

	select {
	case msg1 := <-c1:
		println("Received at channel 1", msg1)

	case msg2 := <-c2:
		println("Received at channel 2", msg2)

	case <-time.After(time.Second * 3):
		println("timeout")

	default:
		println("Default")
	}
}
