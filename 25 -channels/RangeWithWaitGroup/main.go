package main

import "sync"

func main() {
	channel := make(chan int)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(10)
	go publisher(channel)
	go subscriber(channel, &waitGroup)
	waitGroup.Wait()

}

func publisher(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}

func subscriber(channel chan int, waitGroup *sync.WaitGroup) {
	for val := range channel {
		println(val)
		waitGroup.Done()
	}
}
