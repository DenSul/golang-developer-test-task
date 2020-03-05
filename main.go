package main

import (
	configLoader "app/config"
	"app/webserver"
	"log"
)

func main() {

	var config = configLoader.Load()

	log.Println("redis hostname: ", config.Redis.Hostname)

	log.Println("Start web server")
	log.Println(config.WebServer.Port)
	webserver.Start(config.WebServer)
}
