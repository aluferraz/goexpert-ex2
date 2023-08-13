package consulta

import (
	"fmt"
	"github.com/aluferraz/goexpert-ex2/models"
)

type OutputAdapter interface {
	OutputResult(result string) models.HttpOutput
}

type ConsoleOutputAdapter struct {
}

func (c *ConsoleOutputAdapter) OutputResult(result string) models.HttpOutput {
	fmt.Println(result)
	return models.HttpOutput{
		Result: result,
		Error:  nil,
	}
}
