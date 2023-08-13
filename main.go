package main

import (
	"context"
	"github.com/aluferraz/goexpert-ex2/consulta"
	"github.com/aluferraz/goexpert-ex2/models"
	"time"
)

func main() {
	api_cep_channel := make(chan models.HttpOutput)
	via_cep_channel := make(chan models.HttpOutput)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var out consulta.ConsoleOutputAdapter
	var cep = "04538-132"

	apiCep := consulta.ApiCep{
		Cep: cep,
	}
	viaCep := consulta.ViaCep{
		Cep: cep,
	}
	var runner consulta.ConsultaCep

	go runner.Run(&ctx, &apiCep, &out, &api_cep_channel)
	go runner.Run(&ctx, &viaCep, &out, &via_cep_channel)

	select {
	case apiCepResult := <-api_cep_channel:
		if apiCepResult.Error != nil {
			//API Cep executou primeiro, cancelar context para não executar via cep
			cancel()
		}
	case viaCepResult := <-via_cep_channel:
		if viaCepResult.Error != nil {
			//Via Cep executou primeiro, cancelar context para não executar api cep
			cancel()
		}
	case <-ctx.Done():
		out.OutputResult("Timeout reached")
	}

}
