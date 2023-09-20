package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

func addToCache(ctx context.Context, data map[string]interface{}) error {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "PASSWORD",
		DB:       0,
	})

	jsonString, err := json.Marshal(data)

	if err != nil {
		return err
	}

	// set the expiration time to 30 seconds
	err = redisClient.Set("products_cache", jsonString, 30*time.Second).Err()

	if err != nil {
		return nil
	}

	return nil
}
