package main

import "time"

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	go func() {
		channel1 <- 1
	}()
	go publisher(channel1)
	go publisher(channel2)

	select {
	case val := <-channel1:
		println("Channel 1 value: ", val)
	case val := <-channel2:
		println("Channel 2 value: ", val)

	case <-time.After(2 * time.Second):
		println("Timeout")

	default:
		println("No value received")
	}

}

func publisher(channel chan<- int) {
	time.Sleep(3 * time.Second)
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}
