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

type brasilAPIResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type brasilAPIGateway struct{}

func (v *brasilAPIGateway) GetName() string {
	return "BrasilAPI"
}

func (v *brasilAPIGateway) FindAddressByCEP(cep string) (domain.Address, error) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	resp, err := http.Get(url)
	if err != nil {
		return domain.Address{}, err
	}
	if resp.StatusCode != 200 {
		return domain.Address{}, errors.New("[BrasilAPI] request status code is not 200")
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.Address{}, err
	}
	var brasilAPIResponse brasilAPIResponse
	if err = json.Unmarshal(bodyBytes, &brasilAPIResponse); err != nil {
		return domain.Address{}, err
	}

	return domain.Address{
		Cep:        brasilAPIResponse.Cep,
		Logradouro: brasilAPIResponse.Street,
		Bairro:     brasilAPIResponse.Neighborhood,
		Cidade:     brasilAPIResponse.City,
	}, nil
}

func NewBrasilAPICEPGateway() igateway.ICEPGateway {
	return &brasilAPIGateway{}
}
