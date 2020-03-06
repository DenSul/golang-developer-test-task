package providers

import (
	"app/models"
	"app/providers/fileProvider"
)

type Provider interface {
	Load() []models.ParkingTaxi
}

// CreateProvider is a Factory Method
func StrategyFactoryProvider(path string) Provider {
	return &fileProvider.FileProvider{Path: path}
}
