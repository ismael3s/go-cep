package decorators

import (
	"context"

	"github.com/ismael3s/go-cep/internal/application/gateways"
	"github.com/ismael3s/go-cep/internal/application/usecases"
	"github.com/ismael3s/go-cep/internal/domain"
)

type findAddressByCEPDecorator struct {
	useCase      usecases.IFindAddressByCEPUseCase
	cacheGateway gateways.ICacheGateway
}

func (d *findAddressByCEPDecorator) Do(input usecases.FindAddressByCEPInput) (usecases.FindAddressByCEPOutput, error) {
	cep, err := domain.NewCEP(input.Value)
	if err != nil {
		return usecases.FindAddressByCEPOutput{}, err
	}
	cepAddress, _ := d.cacheGateway.Retrieve(context.Background(), cep.GetValue())
	if cepAddress.Cep != "" {
		return usecases.FindAddressByCEPOutput{Address: cepAddress}, nil
	}
	result, err := d.useCase.Do(input)
	if err == nil {
		d.cacheGateway.Persist(context.Background(), result.Address)
	}
	return result, err
}

func NewFindAddressByCEPDecorator(cacheGateway gateways.ICacheGateway, useCase usecases.IFindAddressByCEPUseCase) usecases.IFindAddressByCEPUseCase {
	return &findAddressByCEPDecorator{useCase: useCase, cacheGateway: cacheGateway}
}
