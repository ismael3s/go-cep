package decorators

import (
	"log/slog"

	"github.com/ismael3s/go-cep/internal/application/gateways"
	"github.com/ismael3s/go-cep/internal/domain"
)

type CepGatewayWithLog struct {
	cepGateway gateways.ICEPGateway
}

func NewCepGatewayWithLogging(cepGateway gateways.ICEPGateway) gateways.ICEPGateway {
	return &CepGatewayWithLog{cepGateway: cepGateway}
}

func (d *CepGatewayWithLog) FindAddressByCEP(cep string) (domain.Address, error) {
	address, err := d.cepGateway.FindAddressByCEP(cep)
	if err != nil {
		slog.Error("Error on find address by CEP", "gateway", d.cepGateway.GetName(), "error", err.Error())
	}
	return address, err
}

func (d *CepGatewayWithLog) GetName() string {
	return d.cepGateway.GetName()
}
