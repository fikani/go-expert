package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}

	size, err := f.Write([]byte("Hello, World!"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Wrote %d bytes\n", size)
	f.Close()

	f, err = os.Open("test.txt")
	if err != nil {
		panic(err)
	}

	b := make([]byte, 3)
	for {
		size, err := f.Read(b)
		if size == 0 {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("Read %d bytes: %s\n", size, string(b))
	}

}
