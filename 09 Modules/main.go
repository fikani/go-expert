package main

import (
	"fikani/calc"
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Printf("Sum of 1 and 2 is %f\n", calc.Sum(1.0, 2))
	fmt.Println("a UUID", uuid.New())
}
