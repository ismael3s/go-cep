package usecases

import (
	"errors"
	"time"

	"github.com/ismael3s/go-cep/internal/application/gateways"
	"github.com/ismael3s/go-cep/internal/domain"
)

type FindAddressByCEPUseCase struct {
	cepsGateways []gateways.ICEPGateway
	duration     time.Duration
}

type IFindAddressByCEPUseCase interface {
	Do(input FindAddressByCEPInput) (FindAddressByCEPOutput, error)
}

func defaultDuration(duration time.Duration) time.Duration {
	if duration == 0 {
		return 2 * time.Second
	}
	return duration
}

func NewFindAddresByCEPUseCase(duration time.Duration, cepsGateways ...gateways.ICEPGateway) IFindAddressByCEPUseCase {
	duration = defaultDuration(duration)
	return &FindAddressByCEPUseCase{cepsGateways: cepsGateways, duration: duration}
}

type FindAddressByCEPOutput struct {
	Address domain.Address
}

type FindAddressByCEPInput struct {
	Value string
}

func (u *FindAddressByCEPUseCase) Do(input FindAddressByCEPInput) (FindAddressByCEPOutput, error) {
	cep, err := domain.NewCEP(input.Value)
	if err != nil {
		return FindAddressByCEPOutput{}, err
	}
	addressChan := make(chan *domain.Address)
	var address *domain.Address
	tickerChan := time.NewTicker(u.duration)
	for _, cepGateway := range u.cepsGateways {
		go u.callGateway(cepGateway, cep, addressChan)
	}
	select {
	case address = <-addressChan:
		address.Cep = cep.GetValue()
		return FindAddressByCEPOutput{Address: *address}, nil
	case <-tickerChan.C:
		return FindAddressByCEPOutput{}, errors.New("timeout")
	}

}

func (u *FindAddressByCEPUseCase) callGateway(cepGateway gateways.ICEPGateway, cep *domain.CEP, addressChan chan *domain.Address) {
	address, err := cepGateway.FindAddressByCEP(cep.GetValue())
	if err != nil {
		return
	}
	addressChan <- &address
}
