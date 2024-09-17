package main

import "fmt"

func main() {
	var myMap map[string]int = make(map[string]int)
	mySecondMap := map[string]int{"three": 3, "four": 4}

	myMap["one"] = 1
	myMap["two"] = 2
	delete(mySecondMap, "three")

	fmt.Printf("my map: %v\n", myMap)
	fmt.Printf("my second map: %v\n", mySecondMap)

	for k, v := range myMap {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}
}
