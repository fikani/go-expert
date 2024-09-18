package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"serverside/currency"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", CotacaoHandler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}

func CotacaoHandler(w http.ResponseWriter, r *http.Request) {
	cotacao, err := currency.GetFxRate(time.Millisecond * 200)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cotacaoRepo, close := currency.NewCotacaoRepository()
	defer close()
	err = cotacaoRepo.Save(time.Millisecond*10, cotacao)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(cotacao)
}
