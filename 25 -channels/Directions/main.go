package main

func main() {
	hello := make(chan string)
	go receive("Gophers", hello)
	read(hello)

}

func receive(name string, hello chan<- string) {
	hello <- name
}

func read(hello <-chan string) {
	println("Hello " + <-hello + "!")
}
