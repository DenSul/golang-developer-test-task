package storage

import (
	"app/config"
	"app/models"
	"app/storage/redis"
)

type Storage interface {
	Insert(data []models.ParkingTaxi)
}

func GetStorage(storage config.Storage) Storage {
	// Считай, что это тоже шаблонный метод. :)
	return redis.Connect(storage.Redis)
}
