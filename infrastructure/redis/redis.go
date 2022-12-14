package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/Fajar-Islami/ais_code_test/infrastructure/container"
	"github.com/fatih/color"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(cont container.Redis) *redis.Client {
	redisHost := cont.RedisAddr
	ctx := context.Background()

	if redisHost == "" {
		redisHost = ":6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr:         redisHost,
		MinIdleConns: cont.MinIdleConns,
		PoolSize:     cont.PoolSize,
		PoolTimeout:  time.Duration(cont.PoolTimeout) * time.Second,
		Password:     cont.RedisPassword, // no password set
		DB:           cont.RedisDB,       // use default DB
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic("Failed to create a connection to redis")
	}
	fmt.Println(pong, err)

	color.Green("⇨ Redis status is connected")

	return client
}
