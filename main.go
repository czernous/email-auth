package main

import (
	router "email-auth/router"
	"log"
	"net/http"
)

func main() {

	server := router.SetupRoutes()

	log.Fatal(http.ListenAndServe(":6000", server))

}
