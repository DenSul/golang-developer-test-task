package server

import (
	"app/config"
	"github.com/thedevsaddam/renderer"
	"log"
	"net/http"
)

var rnd *renderer.Render

func Start(config config.WebServer) {
	opts := renderer.Options{
		ParseGlobPattern: config.Tpl + config.TplPattern,
	}

	rnd = renderer.New(opts)
	fs := http.FileServer(http.Dir(config.StaticPath))
	mux := http.NewServeMux()

	setupRoutes(mux)
	mux.Handle(config.StaticPrefix, http.StripPrefix(config.StaticPrefix, fs))

	log.Println("Listening on port ", config.Port)
	err := http.ListenAndServe(config.Port, mux)

	if err != nil {
		log.Fatal(err)
	}
}
