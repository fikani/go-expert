package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Cotacao struct {
	USDBRL struct {
		Bid string `json:"bid"`
	}
}

func main() {
	timeout, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	client := http.Client{}

	req, err := http.NewRequestWithContext(timeout, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Println("Erro ao criar requisição:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao enviar requisição:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Erro ao obter cotação:", resp.Status)
		return
	}

	var result Cotacao
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Erro ao decodificar resposta:", err)
		return
	}

	cotacao := fmt.Sprintf("Dólar: %s", result.USDBRL.Bid)
	fmt.Println(cotacao)

	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte(cotacao))
	if err != nil {
		fmt.Println("Erro ao salvar cotação em arquivo:", err)
	}
}
