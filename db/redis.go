package db

import (
	"context"
	"log/slog"
	"os"

	"github.com/redis/go-redis/v9"
)

var client *RedisClient

type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
	Addr   string
}

var rdb *redis.Client

func NewRedisClient() *RedisClient {
	if client != nil {
		return client
	}
	return &RedisClient{
		Client: rdb,
		Ctx:    context.Background(),
	}
}

func InitRedis() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	slog.Info("Redis", "Connecting", redisHost+":"+redisPort)
	var ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: "",
		DB:       0,
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		slog.Error("Redis", "Error", err)
		panic(err)
	}
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	slog.Info("Redis", "Connected", pong)
}
