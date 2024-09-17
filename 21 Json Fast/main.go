package main

import (
	"fmt"

	"github.com/valyala/fastjson"
)

func main() {

	var parser fastjson.Parser

	jsonData := `{ "name": "John", "age": 30, "city": "New York", "array": [1, 2, 3], "object": { "key": 1 } }`

	jsonObject, err := parser.Parse(jsonData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("NAME=%s\n", jsonObject.GetStringBytes("name"))
	array := jsonObject.GetArray("array")
	for i, value := range array {
		fmt.Printf("ARRAY[%d]=%s\n", i, value)
	}

	object := jsonObject.GetObject("object")
	fmt.Printf("OBJECT.KEY=%v\n", object.Get("key"))

}
