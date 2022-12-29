package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
	zmiddleware "github.com/z3orc/dynamic-rpc/internal/http/middleware"
	"github.com/z3orc/dynamic-rpc/internal/http/routes"
	"github.com/z3orc/dynamic-rpc/internal/util"
)

var port string = util.GetPort()

func main() {

	//ASCII-banner on launch
	util.Banner("CompassAPI")

	//Init router
	router := chi.NewRouter()

	//Middleware
	router.Use(zmiddleware.Recover)
	router.Use(zmiddleware.Logger)
	router.Use(httprate.LimitByIP(
		60,
		60*time.Second,
	))

	//Routes
	routes.Init(router)
	
	//Init listener
	log.Print("| Server listening on ", port, " 🚀")
	log.Fatal(http.ListenAndServe(port, router))
}