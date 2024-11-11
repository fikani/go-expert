package main

import (
	"errors"
	"io"
	"net/http"
	"time"
)

const (
	RequestTimeout = 1 * time.Second
)

func main() {
	address, err := getAddressByZipCode("01001000")
	if err != nil {
		panic(err)
	}
	println(address)
}

// Retrieve address by ZIP code using multiple APIs
func getAddressByZipCode(zipCode string) (string, error) {
	addressChannel := make(chan string)
	go sendResponseToChannel(BrasilAPIURL(zipCode), addressChannel)
	go sendResponseToChannel(ViaCepURL(zipCode), addressChannel)
	defer close(addressChannel)

	select {
	case address := <-addressChannel:
		return address, nil
	case <-time.After(RequestTimeout):
		return "", errors.New("operation timed out after 1 second")
	}
}

// Send API response to channel
func sendResponseToChannel(url string, addressChannel chan string) {
	data, err := makeRequest(url)
	if err == nil {
		addressChannel <- data
	}
}

// Make HTTP GET request and return response body
func makeRequest(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func ViaCepURL(zipCode string) string    { return "https://viacep.com.br/ws/" + zipCode + "/json/" }
func BrasilAPIURL(zipCode string) string { return "https://brasilapi.com.br/api/cep/v1/" + zipCode }
