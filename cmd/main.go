package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/z3orc/dynamic-rpc/handler"
	"github.com/z3orc/dynamic-rpc/middleware"
	"github.com/z3orc/dynamic-rpc/util"
)

var port string = util.GetPort()

func main() {
	router := mux.NewRouter()

	//Middleware
	router.Use(middleware.Recover)
	router.Use(middleware.Logger)

	//Routes
	s := router.PathPrefix("/vanilla").Subrouter()
	s.HandleFunc("/{id}", handler.Vanilla)
	s.HandleFunc("/{id}/download", handler.Vanilla)

	s = router.PathPrefix("/paper").Subrouter()
	s.HandleFunc("/{id}", handler.Paper)
	s.HandleFunc("/{id}/download", handler.Paper)

	s = router.PathPrefix("/purpur").Subrouter()
	s.HandleFunc("/{id}", handler.Purpur)
	s.HandleFunc("/{id}/download", handler.Purpur)

	//Static index
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	
	//ASCII-banner on launch
	util.Banner()
	log.Fatal(http.ListenAndServe(port, router))
}