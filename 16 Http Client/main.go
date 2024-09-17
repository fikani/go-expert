package main

import (
	"io"
	"net/http"
	"strings"
	"time"
)

func main() {
	client := http.Client{Timeout: time.Duration(4) * time.Second}
	res, err := client.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))

	jsonData := `{"key": "value"}`
	res, err = client.Post("https://httpbin.org/post", "text/plain", strings.NewReader(jsonData))
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err = io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))

	req, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent", "my-client")
	res, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err = io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
