package main

import "time"

func main() {
	data := make(chan int)
	qtyWorkers := 100

	for i := 0; i < qtyWorkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 100; i++ {
		data <- i
	}

	close(data)
	time.Sleep(time.Second)

}

func worker(workerId int, data chan int) {
	for val := range data {
		println("worker", workerId, "received", val)
		time.Sleep(time.Second)
	}
}
