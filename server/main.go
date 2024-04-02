package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sheeiavellie/mpm-coursework/server/handlers"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		return
	}

	port := os.Getenv("PORT")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /process", handlers.HandleProcess)

	log.Printf("seuserrver is listening on port: %s", port)
	http.ListenAndServe(":"+port, mux)
}
