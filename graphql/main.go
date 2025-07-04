package graphql

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	accountURL string `envconfig:"ACCOUNT_SERVICE_URL" default:"http://localhost:8080"`
	productURL string `envconfig:"PRODUCT_SERVICE_URL" default:"http://localhost:8081"`
	orderURL   string `envconfig:"ORDER_SERVICE_URL" default:"http://localhost:8082"`
}

func main() {
	var cfg AppConfig
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	s, err := NewGraphQLServer(cfg.accountURL, cfg.productURL, cfg.orderURL)
	if err != nil {
		log.Fatalf("Failed to create GraphQL server: %v", err)
	}
	log.Println("GraphQL server started with the following configuration:")
	log.Printf("Account Service URL: %s", cfg.accountURL)
	log.Printf("Product Service URL: %s", cfg.productURL)
	log.Printf("Order Service URL: %s", cfg.orderURL)
	srv := handler.New(s.ToExecutableSchema())
	http.Handle("/graphql", srv)
	http.Handle("/playground", playground.Handler("GraphQL Playground", "/graphql"))

	log.Fatal(http.ListenAndServe(":8083", nil))
}
