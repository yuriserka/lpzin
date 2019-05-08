package routes

import (
	"flag"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/yuriserka/lpzin/controllers"
)

var router *mux.Router

func init() {
	var dir string
	flag.StringVar(&dir, "dir", "./static/", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	router = mux.NewRouter()
	router.HandleFunc("/", controllers.MainHandler)
	router.HandleFunc("/lindo", controllers.Handler)
	// assim n funfa os arquivos de CSS
	// router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(dir)))
}

func Run() {
	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
