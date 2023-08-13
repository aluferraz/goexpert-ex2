package models

type ViaCepOutput struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type ApiCepOutput struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

//Normalizando output do projeto

type CepOutput struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Cidade     string `json:"cidade"`
	Uf         string `json:"uf"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

type HttpOutput struct {
	Result string
	Error  error
}
