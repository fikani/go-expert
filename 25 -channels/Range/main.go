package main

func main() {
	channel := make(chan int)
	go publisher(channel)
	subscriber(channel)

}

func publisher(channel chan int) {
	for i := 0; i < 10; i++ {
		channel <- i
	}
	close(channel)
}

func subscriber(channel chan int) {
	for val := range channel {
		println(val)
	}
}
