package consulta

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aluferraz/goexpert-ex2/models"
	"io"
	"net/http"
)

type ConsultaCepInterface interface {
	GetApiUrl() string
	MapResponse(body []byte) (models.CepOutput, error)
}

type ConsultaCep struct {
	adapter OutputAdapter
	url     string
}

func (c ConsultaCep) outputResult(result string) models.HttpOutput {
	return c.adapter.OutputResult(result)
}

func (c ConsultaCep) handleError(err error) models.HttpOutput {
	var result models.CepOutput
	result.Ok = false
	result.StatusText = err.Error()

	resultText, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	return c.outputResult(fmt.Sprintf("[ERROR] '%s' - API consultada '%s' ", string(resultText), c.url))
}

func (c ConsultaCep) Run(ctx *context.Context, cepinterface ConsultaCepInterface, adapter OutputAdapter, channel *chan models.HttpOutput) models.HttpOutput {
	c.adapter = adapter
	c.url = cepinterface.GetApiUrl()
	req, err := http.NewRequestWithContext(*ctx, "GET", c.url, nil)
	if err != nil {
		return c.handleError(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return c.handleError(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body) // response body is []byte
	if err != nil {
		return c.handleError(err)
	}
	result, err := cepinterface.MapResponse(body)
	if err != nil {
		return c.handleError(err)
	}
	resultJson, err := json.Marshal(result)
	if err != nil {
		return c.handleError(err)
	}
	outputText := fmt.Sprintf("API consultada '%s', resposta '%s'", c.url, string(resultJson))

	httpOutput := c.outputResult(outputText)
	if channel != nil {
		*channel <- httpOutput
	}
	return httpOutput
}
