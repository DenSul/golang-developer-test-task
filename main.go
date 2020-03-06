package main

import (
	"app/cmd"
	configLoader "app/config"
	"app/providers"
	"app/server"
	dataLoader "app/services/loader"
	"app/storage"
	"log"
)

func main() {

	var args = new(cmd.Args)
	args.Parse()

	var config = configLoader.Load(args)
	var storage = storage.GetStorage(config.Storage)

	if args.Source != "" {
		var provider = providers.StrategyFactoryProvider(args.Source)
		var dataLoader = dataLoader.Loader{
			Storage:  storage,
			Provider: provider,
		}
		dataLoader.Run()
	}

	log.Println("Start web server")
	server.Start(config.WebServer)
}
