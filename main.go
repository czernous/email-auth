package main

import (
	router "email-auth/router"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func tryGetEnvVars(vars [7]string) {
	// iterate
	for _, v := range vars {
		temp := os.Getenv(v)
		if len(temp) < 1 {
			log.Fatalf("Could not load variable %s", v)
			os.Exit(1)
		}
	}
}

func main() {

	err := godotenv.Load()

	port := os.Getenv("AUTH_API_PORT")

	tryGetEnvVars([7]string{
		"AUTH_JWT_SECRET",
		"AUTH_API_PORT",
		"AUTH_API_KEY",
		"SMTP_HOST",
		"SMTP_PORT",
		"SMTP_LOGIN",
		"SMTP_PASSWORD"})

	if err != nil {
		log.Printf("Error loading .env file. Err: %s", err)
	}

	server := router.SetupRoutes()

	log.Fatal(http.ListenAndServe(":"+port, cors.AllowAll().Handler(server)))

}
