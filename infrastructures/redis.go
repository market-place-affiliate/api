package infrastructure

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func NewRedis(host string, port int, db int, username string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		DB:       db,
		Username: username,
		Password: password,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("failed to ping redis: %s\n", err.Error())
	}
	return client
}
