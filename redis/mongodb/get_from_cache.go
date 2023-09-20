package main

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis"
)

func getFromCache(ctx context.Context) (bool, map[string]interface{}, error) {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "PASSWORD",
		DB:       0,
	})

	productsCache, err := redisClient.Get("products_cache").Bytes()

	if err != nil {
		return false, nil, nil
	}

	res := map[string]interface{}{}

	err = json.Unmarshal(productsCache, &res)

	if err != nil {
		return false, nil, nil
	}

	return true, res, nil
}
