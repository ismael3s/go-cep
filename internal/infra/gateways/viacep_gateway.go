package gateways

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	igateway "github.com/ismael3s/go-cep/internal/application/gateways"
	"github.com/ismael3s/go-cep/internal/domain"
)

type viaCEPResponse struct {
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

type viaCEPGateway struct{}

func (v *viaCEPGateway) GetName() string {
	return "ViaCEP"
}

func (v *viaCEPGateway) FindAddressByCEP(cep string) (domain.Address, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return domain.Address{}, err
	}
	if resp.StatusCode != 200 {
		return domain.Address{}, errors.New("[ViaCEP] request status code is not 200")
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.Address{}, err
	}
	var viaCEPResp viaCEPResponse
	if err = json.Unmarshal(bodyBytes, &viaCEPResp); err != nil {
		return domain.Address{}, err
	}

	return domain.Address{
		Cep:        viaCEPResp.Cep,
		Logradouro: viaCEPResp.Logradouro,
		Bairro:     viaCEPResp.Bairro,
		Cidade:     viaCEPResp.Localidade,
	}, nil
}

func NewViaCEPGateway() igateway.ICEPGateway {
	return &viaCEPGateway{}
}
