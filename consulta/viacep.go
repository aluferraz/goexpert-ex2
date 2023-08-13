package consulta

import (
	"encoding/json"
	"fmt"
	"github.com/aluferraz/goexpert-ex2/models"
)

type ViaCep struct {
	Cep string
}

func (v ViaCep) GetApiUrl() string {
	return fmt.Sprintf("https://viacep.com.br/ws/%s/json", v.Cep)
}
func (v ViaCep) MapResponse(body []byte) (models.CepOutput, error) {
	var viaCepReponse models.ViaCepOutput
	err := json.Unmarshal(body, &viaCepReponse)
	if err != nil {
		return models.CepOutput{}, err
	}
	return models.CepOutput{
		Cep:        viaCepReponse.Cep,
		Logradouro: viaCepReponse.Logradouro,
		Cidade:     viaCepReponse.Localidade,
		Uf:         viaCepReponse.Uf,
	}, nil

}
