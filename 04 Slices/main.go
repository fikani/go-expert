package main

import "fmt"

func main() {

	slice := []int{1, 34, 53, 23, 12, 45, 67, 78, 89, 90}
	fmt.Printf("O slice tem cap de %d e len de %d\n", cap(slice), len(slice))
	fmt.Printf("O slice tem cap de %d e len de %d\n", cap(slice[:0]), len(slice[:0]))
	fmt.Printf("O slice tem cap de %d e len de %d\n", cap(slice[:4]), len(slice[:4]))
	fmt.Printf("O slice tem cap de %d e len de %d, array=%v\n", cap(slice[2:]), len(slice[2:]), slice[2:])

	slice = append(slice, 100)
	fmt.Printf("O slice tem cap de %d e len de %d\n", cap(slice), len(slice))
}
