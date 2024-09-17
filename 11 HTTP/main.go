package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	defer fmt.Println("Primeiro")
	defer fmt.Println("Segundo")
	fmt.Println("Terceiro")
}
