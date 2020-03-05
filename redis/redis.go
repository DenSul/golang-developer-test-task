package redis

import (
	"app/config"
	"fmt"
	"github.com/go-redis/redis/v7"
	"log"
	"strconv"
)

var client *redis.Client

func Connect(config config.RedisConfiguration) *redis.Client {

	client = redis.NewClient(&redis.Options{
		Addr:     config.Hostname + ":" + strconv.Itoa(config.Port),
		Password: config.Password,
		DB:       config.Database,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		log.Fatal("i can not connected to Redis. Check config file.", err)
	}

	fmt.Println(pong)

	return client
}

func GetDB() *redis.Client {
	return client
}
