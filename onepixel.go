package main

import (
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/didiercrunch/onepixel/data"
	"github.com/gorilla/mux"
)

const STATIC_DIR_PATH = "static/"

func serveStatic(router *mux.Router) {
	handler := func(w http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		filepath := "/" + vars["path"]
		w.Header().Set("Cache-Control", "public, max-age=43200")
		http.ServeFile(w, request, path.Join(STATIC_DIR_PATH, filepath))
	}
	router.HandleFunc("/{path:.*}", handler)
}

func serveApi(router *mux.Router, apiHandler http.Handler) {
	router.Handle("/api/{serie_name}/{url:.*}", apiHandler)
}

func createMuxRouter() http.Handler {
	r := mux.NewRouter()
	serveApi(r, &data.Data{params.InfluxDBClient})
	serveStatic(r.PathPrefix("/").Subrouter())
	return r
}

func main() {
	r := createMuxRouter()
	http.Handle("/", r)
	address := params.Host + ":" + strconv.Itoa(params.Port)
	if err := http.ListenAndServe(address, createMuxRouter()); err != nil {
		log.Fatal(err)
	}
}
