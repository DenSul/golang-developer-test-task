package main

import (
	configLoader "app/config"
	"app/providers"
	"app/redis"
	"app/webserver"
	"log"
)

func main() {

	var config = configLoader.Load()

	log.Println("redis hostname: ", config.Redis.Hostname)

	log.Println("Start web server")
	log.Println(config.WebServer.Port)
	providers.SelectStrategyDownloadFile("dsa")
	redis.Connect(config.Redis)
	webserver.Start(config.WebServer)
}
