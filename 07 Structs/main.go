package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

type Person interface {
	setStreet(street string)
}

// Compisition Client.State works because of the Address struct
type Client struct {
	Name string
	Age  int
	Address
}

// This is the same as the Client struct
type Client02 struct {
	Name    string
	Age     int
	Address Address
}

func NewClient(name string, age int, street string, city string, state string, zipCode string) *Client {
	return &Client{
		Name: name,
		Age:  age,
		Address: Address{
			Street:  street,
			City:    city,
			State:   state,
			ZipCode: zipCode,
		},
	}
}

func (c *Client) setStreet(street string) {
	println(c)
	c.Street = street
}

func main() {
	user01 := Client{
		Name: "John Doe",
		Age:  30,
		Address: Address{
			Street:  "123 Main St",
			City:    "Springfield",
			State:   "IL",
			ZipCode: "62701",
		},
	}

	user01.Age = 31

	userInterface := &user01
	userInterface.setStreet("456 Main St")
	fmt.Printf("The street of user01 is: %v\n", user01.Address.Street)
}
