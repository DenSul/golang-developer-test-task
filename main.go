package main

import (
	configLoader "app/config"
	"app/providers"
	"app/server"
	"app/storage/redis"
	"fmt"
	"log"
)

func main() {

	var config = configLoader.Load()

	log.Println("Start web server")
	log.Println(config.WebServer.Port)

	redis.Connect(config.Redis)

	test := providers.StrategyFactoryProvider("resources/data/data.json")
	fmt.Println(test.Load())
	server.Start(config.WebServer)
}
