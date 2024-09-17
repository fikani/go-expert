package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{Name: "John Doe", Age: 30}

	jsonData, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}

	println(string(jsonData))

	err = json.NewEncoder(os.Stdout).Encode(person)
	if err != nil {
		panic(err)
	}

	jsonData2 := []byte(`{"name":"Jane Doe","age":25}`)
	var person2 Person
	err = json.Unmarshal(jsonData2, &person2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Name: %s, Age: %d\n", person2.Name, person2.Age)

}
