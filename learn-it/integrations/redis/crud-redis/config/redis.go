package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	Ctx = context.Background()
	Rdb *redis.Client
)

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		PoolSize: 10,
		DB:       0,
	})

	/*
		In order to test first redis connection, we will do one Ping-Pong
	*/
	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Please check if redis is running or not using ```ps -ef | grep redis``` command %v", err)
	}
}
