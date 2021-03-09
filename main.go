package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	account "github.com/heru-wijaya/go-graphql-skeleton/account"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	AccountURL string `envconfig:"ACCOUNT_SERVICE_URL"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	accUrl := os.Getenv("ACCOUNT_SERVICE_URL")

	s, err := account.NewGraphQLServer(accUrl)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/graphql", handler.GraphQL(s.ToExecutableSchema()))
	http.Handle("/playground", handler.Playground("Spidey", "/graphql"))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
