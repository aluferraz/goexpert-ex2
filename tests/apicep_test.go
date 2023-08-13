package tests

import (
	"context"
	"github.com/aluferraz/goexpert-ex2/consulta"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestApiCep(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	assert.New(t)

	var out consulta.ConsoleOutputAdapter

	var cep = "04538-132"

	apiCep := consulta.ApiCep{
		Cep: cep,
	}

	var runner consulta.ConsultaCep

	response := runner.Run(&ctx, &apiCep, &out, nil)
	// assert for nil (good for errors)
	assert.Nil(t, response.Error)
	assert.Contains(t, response.Result, "https://cdn.apicep.com")
	assert.Contains(t, response.Result, cep)
}
