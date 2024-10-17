package repositories

import (
	"context"
	"errors"

	"example.com/m/internal/config"
	"github.com/redis/go-redis/v9"
)

type TokenRepository struct {
	rdb *redis.Client
}

func NewTokenRepository(r *redis.Client) *TokenRepository {
	return &TokenRepository{
		rdb: r,
	}
}

func (r *TokenRepository) GetByEmail(ctx *context.Context, email string) (*string, error) {
	v, err := r.rdb.Get(*ctx, email).Result()

	if errors.Is(err, redis.Nil) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &v, nil
}

func (r *TokenRepository) Set(ctx *context.Context, email string, token string) error {
	err := r.rdb.Set(*ctx, email, token, config.Config.JWTExpiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *TokenRepository) DeleteByEmail(ctx *context.Context, email string) error {
	err := r.rdb.Del(*ctx, email).Err()
	if err != nil {
		return err
	}
	return nil
}
