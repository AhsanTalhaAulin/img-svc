package conn

import (
	"context"
	"log"

	"github.com/go-redis/redis/v9"
)

type goRedis struct {
	Rdb *redis.Client
}

var RedisClient goRedis

func ConnectRedis() error {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 15,
	})

	RedisClient.Rdb = rdb

	var ctx = context.Background()

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Println("error in setting key-value pair")
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		log.Println("error in getting key-value pair")
		panic(err)
	}
	log.Println("key", val)
	log.Println("Redis connection established")

	return nil
}
