package redis

import (
	"app/config"
	"app/models"
	"fmt"
	"github.com/go-redis/redis/v7"
	"strconv"
)

var client *redis.Client

type Redis struct {
	Config config.RedisConfiguration
}

func (r Redis) Connect() *redis.Client {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     r.Config.Hostname + ":" + strconv.Itoa(r.Config.Port),
		Password: r.Config.Password,
		DB:       r.Config.Database,
	})

	_, err := redisClient.Ping().Result()

	if err != nil {
		panic(fmt.Sprintf("i can not connected to Redis. Check config file. %s", err))
	}

	client = redisClient

	return redisClient
}

func (r Redis) Insert(data []models.ParkingTaxi) {

}

func (r Redis) GetDB() *redis.Client {
	return client
}
