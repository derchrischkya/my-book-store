package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/derchrischkya/libary/authentication"
	"github.com/derchrischkya/libary/endpoints/healthcheck"
	"github.com/derchrischkya/libary/endpoints/inventory"
	"github.com/julienschmidt/httprouter"
)

func main() {
	//HTTP server for the api
	router := httprouter.New()

	// Technical Routes
	router.GET("/api/ping", authentication.Authenticate(healthcheck.Ping()))

	// Command routes. (verb + noun, imperative)
	router.POST("/api/inventory/capture-book", authentication.Authenticate(inventory.CaptureBook()))
	// router.POST("/api/lending/lend-book", inventory.LendBook())

	// Query routes (noun)
	// router.GET("/api/inventory/books", inventory.Books())

	port := 3000
	address := fmt.Sprintf(":%d", port)

	log.Println("Started: server is running on port", port)
	err := http.ListenAndServe(address, router)
	log.Fatal(err)
}
