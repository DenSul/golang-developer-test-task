package main

import (
	configLoader "app/config"
	"log"
	"net/http"

	"github.com/thedevsaddam/renderer"
)

var rnd *renderer.Render

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./resources/tpl/*.html",
	}

	rnd = renderer.New(opts)
}

func main() {
	var config = configLoader.Load()

	log.Println("redis hostname: ", config.Redis.Hostname)

	fs := http.FileServer(http.Dir("./resources/static"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	port := ":9000"
	log.Println("Listening on port ", port)
	err := http.ListenAndServe(port, mux)

	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "home", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "about", nil)
}
