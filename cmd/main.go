package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/cadugr/desafio-multithreading-go/internal/dto"
	webservices "github.com/cadugr/desafio-multithreading-go/internal/web"
)

func main() {
	//Criando os canais para receber as respostas das goroutines
	totalParametros := len(os.Args)
	if totalParametros == 1 {
		fmt.Println("Para executar o programa você deve informar 1 parâmetro (cep).  Exemplo: go run cmd/main.go 24230050")
		os.Exit(1)
	}
	cep := os.Args[1]
	responseBrasilApiCep := make(chan dto.Response)
	responseViaCep := make(chan dto.Response)

	go func() {
		time.Sleep(time.Second)
		brasilApiCep, _ := webservices.CallBrasilApiCep(cep)
		json, _ := json.Marshal(brasilApiCep)
		responseBrasilApiCep <- dto.Response{Msg: "Resultado fornecido pela API BrasilApiCep: ", BrasilApiCep: string(json)}
	}()

	go func() {
		time.Sleep(time.Second)
		viacep, _ := webservices.CallViaCep(cep)
		json, _ := json.Marshal(viacep)
		responseViaCep <- dto.Response{Msg: "Resultado fornecido pela API ViaCep: ", ViaCep: string(json)}
	}()

	select {
	case responseBrasilApiCep := <-responseBrasilApiCep:
		fmt.Println(responseBrasilApiCep.Msg, responseBrasilApiCep.BrasilApiCep)
	case responseViaCep := <-responseViaCep:
		fmt.Println(responseViaCep.Msg, responseViaCep.ViaCep)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout")
	}

}
