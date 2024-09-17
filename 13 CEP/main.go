package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	for _, cep := range os.Args[1:] {
		res, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		jsonData, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		var endereco Endereco
		err = json.Unmarshal(jsonData, &endereco)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Encontrado - Cep: %s, Logradouro: %s\n", endereco.Cep, endereco.Logradouro)
		f, err := os.Create("endereco_" + cep + ".txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.WriteString("Cep: " + endereco.Cep + "\n")
		f.WriteString("Logradouro: " + endereco.Logradouro + "\n")
		f.WriteString("Complemento: " + endereco.Complemento + "\n")
		f.WriteString("Unidade: " + endereco.Unidade + "\n")
		f.WriteString("Bairro: " + endereco.Bairro + "\n")
		f.WriteString("Localidade: " + endereco.Localidade + "\n")
		f.WriteString("Uf: " + endereco.Uf + "\n")
		f.WriteString("Ibge: " + endereco.Ibge + "\n")
		f.WriteString("Gia: " + endereco.Gia + "\n")
		f.WriteString("Ddd: " + endereco.Ddd + "\n")
		f.WriteString("Siafi: " + endereco.Siafi + "\n")

	}
}
