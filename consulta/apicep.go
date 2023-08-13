package consulta

import (
	"encoding/json"
	"fmt"
	"github.com/aluferraz/goexpert-ex2/models"
)

type ApiCep struct {
	Cep string
}

func (v ApiCep) GetApiUrl() string {
	return fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", v.Cep)
}
func (v ApiCep) MapResponse(body []byte) (models.CepOutput, error) {
	var apiCepReponse models.ApiCepOutput
	err := json.Unmarshal(body, &apiCepReponse)
	if err != nil {
		return models.CepOutput{}, err
	}
	return models.CepOutput{
		Cep:        apiCepReponse.Code,
		Logradouro: apiCepReponse.Address,
		Cidade:     apiCepReponse.City,
		Uf:         apiCepReponse.State,
	}, nil

}
