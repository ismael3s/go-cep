package gateways

import (
	"context"

	"github.com/ismael3s/go-cep/internal/domain"
)

type ICacheGateway interface {
	Persist(ctx context.Context, cepAddress domain.Address) error
	Retrieve(ctx context.Context, cep string) (domain.Address, error)
}
