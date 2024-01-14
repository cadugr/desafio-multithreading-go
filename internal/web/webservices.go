package webservices

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/cadugr/desafio-multithreading-go/internal/dto"
)

func CallViaCep(cep string) (dto.ViaCep, error) {
	resp, error := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return dto.ViaCep{}, error
	}
	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return dto.ViaCep{}, error
	}
	var viaCep dto.ViaCep
	error = json.Unmarshal(body, &viaCep)
	if error != nil {
		return dto.ViaCep{}, error
	}
	return viaCep, nil
}

func CallBrasilApiCep(cep string) (dto.BrasilApiCep, error) {
	resp, error := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if error != nil {
		return dto.BrasilApiCep{}, error
	}
	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return dto.BrasilApiCep{}, error
	}
	var brasilApiCep dto.BrasilApiCep
	error = json.Unmarshal(body, &brasilApiCep)
	if error != nil {
		return dto.BrasilApiCep{}, error
	}
	return brasilApiCep, nil
}
