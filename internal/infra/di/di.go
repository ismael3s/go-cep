package di

import (
	"time"

	"github.com/ismael3s/go-cep/internal/application/decorators"
	usecases "github.com/ismael3s/go-cep/internal/application/usecases"
	gateways "github.com/ismael3s/go-cep/internal/infra/gateways"
)

func FindAddressByCEPDI() usecases.IFindAddressByCEPUseCase {
	viacepGateway := decorators.NewCepGatewayWithLogging(gateways.NewViaCEPGateway())
	brasilAPICEPGateway := decorators.NewCepGatewayWithLogging(gateways.NewBrasilAPICEPGateway())
	duration := 2 * time.Second
	useCase := usecases.NewFindAddresByCEPUseCase(duration, viacepGateway, brasilAPICEPGateway)
	return decorators.NewFindAddressByCEPDecorator(
		gateways.NewRedisCacheGateway(),
		useCase,
	)
}
