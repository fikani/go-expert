package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(1, 2)
	t := 2
	a := 3
	if err != nil {
		fmt.Println(err, a)
	} else {
		fmt.Println(valor, t)
	}

	fmt.Println(sumMultiple(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	// Closures
	total := func() int {
		return 10 * sumMultiple(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}()

	fmt.Println(total)
}

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("a soma é maior ou igual a 50")
	}
	return a + b, nil
}

func sumMultiple(a int, b ...int) int {
	total := a
	for _, v := range b {
		total += v
	}
	return total
}
