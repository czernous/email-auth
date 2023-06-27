package main

import (
	router "email-auth/router"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	secret := os.Getenv("AUTH_JWT_SECRET")

	if len(secret) < 1 {
		log.Fatalf("AUTH_JWT_SECRET is not set")
	}

	if err != nil {
		log.Fatalf("Error loading .env file. Err: %s", err)
	}

	server := router.SetupRoutes()

	log.Fatal(http.ListenAndServe(":5001", server))

}
