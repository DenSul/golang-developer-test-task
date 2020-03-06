package main

import (
	"app/cmd"
	configLoader "app/config"
	"app/providers"
	"app/server"
	"app/storage/redis"
	"fmt"
	"log"
)

func main() {

	var args = new(cmd.Args)
	args.Parse()
	fmt.Println(args)
	var config = configLoader.Load(args)

	log.Println("Start web server")
	log.Println(config.WebServer.Port)

	redis.Connect(config.Redis)

	if args.Source != "" {
		providers.StrategyFactoryProvider(args.Source)
		//fmt.Println(test.Load())
	}

	server.Start(config.WebServer)
}
