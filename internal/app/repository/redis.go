package repository

import (
	"context"
	"errors"
	"time"

	"github.com/cavelms/config"
)

type redisDB struct {
	context.Context
}

type Redis interface {
	Set(key, val string, exp int) error
	Get(key string) (string, error)
}

func newRedisRepository() Redis {
	return &redisDB{
		context.Background(),
	}
}

func (ctx *redisDB) Set(key, val string, exp int) error {
	rdb := config.RedisClient(0)
	defer rdb.Close()

	expireTime := time.Duration(exp) * time.Second
	err := rdb.Set(ctx, key, val, expireTime).Err()
	if err != nil {
		return errors.New("RequestTokens(): rdb.Set: accessToken: " + err.Error())
	}

	return nil
}

func (ctx *redisDB) Get(key string) (string, error) {
	rdb := config.RedisClient(0)
	defer rdb.Close()

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}
