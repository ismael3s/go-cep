package domain

type Address struct {
	Cep        string `json:"cep" redis:"cep"`
	Logradouro string `json:"logradouro" redis:"logradouro"`
	Bairro     string `json:"bairro" redis:"bairro"`
	Cidade     string `json:"cidade" redis:"cidade"`
}
