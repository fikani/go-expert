package currency

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CotacaoDTO struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func GetFxRate(timeout time.Duration) (CotacaoDTO, error) {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutCtx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		fmt.Printf("Erro ao criar requisição: %v\n", err)
		return CotacaoDTO{}, err
	}
	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Erro ao enviar requisição: %v\n", err)
		return CotacaoDTO{}, err
	}

	defer resp.Body.Close()
	var result CotacaoDTO
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Printf("Erro ao decodificar resposta: %v\n", err)
		return CotacaoDTO{}, err
	}

	return result, nil
}
