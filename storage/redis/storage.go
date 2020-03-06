package redis

import (
	"app/config"
	"github.com/go-redis/redis/v7"
	"log"
	"strconv"
)

var client *redis.Client

func Connect(config config.RedisConfiguration) *redis.Client {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Hostname + ":" + strconv.Itoa(config.Port),
		Password: config.Password,
		DB:       config.Database,
	})

	_, err := redisClient.Ping().Result()

	if err != nil {
		log.Fatal("i can not connected to Redis. Check config file.", err)
	}

	client = redisClient

	return redisClient
}

func GetDB() *redis.Client {
	return client
}
