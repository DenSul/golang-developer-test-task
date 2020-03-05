package webserver

import (
	"app/webserver/web"
	"net/http"
)

func setupRoutes(mux *http.ServeMux) {
	// Inject render html
	web.Rnd = rnd

	mux.HandleFunc("/", web.Home)
	mux.HandleFunc("/about", about)
}

func about(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "about", nil)
}
