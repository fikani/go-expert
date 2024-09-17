package main

import "fmt"

type MyInt int

type Number interface {
	~int | float64
}

func Sum[T Number](a, b T) T {
	return a + b
}

func Equal[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	var i1 MyInt = 1
	var i2 MyInt = 2
	fmt.Printf("Sum of 1 and 2 is %d\n", Sum(i1, i2))
	fmt.Printf("Sum of 1 and 2 is %f\n", Sum(1.0, 2))
	fmt.Printf("Are 1 and 2 equal? %t\n", Equal(1, 2))
	fmt.Printf("Are 1 and 1 equal? %t\n", Equal(100, 1.3))
}
