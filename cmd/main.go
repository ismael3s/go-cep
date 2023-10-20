package main

import (
	"log"
	"os"

	"github.com/ismael3s/go-cep/internal/infra/rest"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	log.Fatalln(
		rest.SetupRestServer().Run(port),
	)
}
