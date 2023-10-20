package gateways

import (
	"context"
	"log"
	"os"

	igateway "github.com/ismael3s/go-cep/internal/application/gateways"
	"github.com/ismael3s/go-cep/internal/domain"
	"github.com/redis/go-redis/v9"
)

type redisCacheGateway struct {
	redisClient *redis.Client
}

func (r *redisCacheGateway) Persist(ctx context.Context, cepAddress domain.Address) error {
	_, err := r.redisClient.HSet(ctx, cepAddress.Cep, cepAddress).Result()
	return err
}

func (r *redisCacheGateway) Retrieve(ctx context.Context, cep string) (domain.Address, error) {
	result, err := r.redisClient.HGetAll(ctx, cep).Result()
	if err != nil {
		if err == redis.Nil {
			log.Println("key does not exist")
			return domain.Address{}, nil
		}
		return domain.Address{}, err
	}
	return domain.Address{
		Cep:        result["cep"],
		Logradouro: result["logradouro"],
		Bairro:     result["bairro"],
		Cidade:     result["cidade"],
	}, nil
}

func NewRedisCacheGateway() igateway.ICacheGateway {
	redisDSN := os.Getenv("REDIS_URL")
	if redisDSN == "" {
		redisDSN = "redis://:root@localhost:6379"
	}
	redisOptions, err := redis.ParseURL(redisDSN)
	if err != nil {
		log.Fatal(err)
	}
	redisClient := redis.NewClient(redisOptions)
	return &redisCacheGateway{redisClient: redisClient}
}
