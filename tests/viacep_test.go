package tests

import (
	"context"
	"github.com/aluferraz/goexpert-ex2/consulta"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestViaCep(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	assert.New(t)

	var out consulta.ConsoleOutputAdapter

	var cep = "04538-132"

	viaCep := consulta.ViaCep{
		Cep: cep,
	}

	var runner consulta.ConsultaCep

	response := runner.Run(&ctx, &viaCep, &out, nil)
	// assert for nil (good for errors)
	assert.Nil(t, response.Error)
	assert.Contains(t, response.Result, "https://viacep.com.br")
	assert.Contains(t, response.Result, cep)
}
