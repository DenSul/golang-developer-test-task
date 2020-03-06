package storage

import (
	"app/config"
	"app/models"
	"app/storage/redis"
)

type Storage interface {
	//Connect()
	//GetDB()
	Insert(data []models.ParkingTaxi)
	//CreateIndex(indexName string, data []string)
	//FindBy(keyName string, query string)
}

func GetStorage(storage config.Storage) Storage {
	return &redis.Redis{Config: storage.Redis}
}
