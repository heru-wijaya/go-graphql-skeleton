package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	account "github.com/heru-wijaya/go-graphql-skeleton/account"
	dummy "github.com/heru-wijaya/go-graphql-skeleton/dummy"
	"github.com/joho/godotenv"
)

// AppConfig struct type for env variable
type AppConfig struct {
	AccountURL string `envconfig:"ACCOUNT_SERVICE_URL"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	accURL := os.Getenv("ACCOUNT_SERVICE_URL")

	s, err := account.NewGraphQLServer(accURL)
	if err != nil {
		log.Fatal(err)
	}

	d, err := dummy.NewGraphQLServer()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/account", handler.GraphQL(s.ToExecutableSchema()))
	http.Handle("/dummy", handler.GraphQL(d.ToExecutableSchema()))
	http.Handle("/playground", handler.Playground("account", "/dummy"))

	log.Fatal(http.ListenAndServe(":8081", nil))
}
