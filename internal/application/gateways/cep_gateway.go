package gateways

import "github.com/ismael3s/go-cep/internal/domain"

type ICEPGateway interface {
	FindAddressByCEP(string) (domain.Address, error)
	GetName() string
}
