package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Endereco struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/afif", Afif{name: "Afif2"})
	mux.HandleFunc("/", BuscarCepHandler)
	http.ListenAndServe(":8080", mux)
}

type Afif struct {
	name string
}

func (a Afif) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello, %v!", a.name)))
}

func BuscarCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cep" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	encoder := json.NewEncoder(w)
	endereco, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder.Encode(endereco)
}

func BuscaCep(cep string) (*Endereco, error) {
	res, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var endereco Endereco
	err = json.Unmarshal(jsonData, &endereco)
	if err != nil {
		return nil, err
	}

	return &endereco, nil
}
