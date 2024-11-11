package main

import "fmt"

// Thread 1
func main() {
	// Code here
	channel := make(chan string)

	go func() {
		channel <- "Hello, World!"
		fmt.Println("Message sent to channel")
		channel <- "Hello, World 2!"
		fmt.Println("Message sent to channel")
	}()

	msg := <-channel
	msg2 := <-channel
	fmt.Println(msg)
	fmt.Println(msg2)
}
