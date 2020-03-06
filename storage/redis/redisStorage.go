package redis

import (
	"app/config"
	"app/models"
	"app/storage"
	"fmt"
	"github.com/go-redis/redis/v7"
	"strconv"
)

type Redis struct {
	Client *redis.Client
}

func Connect(config config.RedisConfiguration) storage.Storage {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Hostname + ":" + strconv.Itoa(config.Port),
		Password: config.Password,
		DB:       config.Database,
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic(fmt.Sprintf("i can not connected to Redis. Check config file. %s", err))
	}

	return &Redis{Client: client}
}

func (r Redis) Insert(data []models.ParkingTaxi) {

}

func (r Redis) GetDB() *redis.Client {
	return r.Client
}
