package main

import "fmt"

func main() {
	meuArray := [5]int{1, 2, 3, 4, 5}

	println(meuArray[0])
	println(meuArray[len(meuArray)-1])

	for i, v := range meuArray {
		fmt.Printf("índice: %d, valor: %d\n", i, v)
	}
}
