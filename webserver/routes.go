package webserver

import "net/http"

func handleRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
}

func home(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "home", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "about", nil)
}
