package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq" // PostgreSQL driver

	account "github.com/Aadithya-J/microservices-ecommerce/account"
)

// adjust import path as needed

func main() {
	log.Println("Account service starting...")

	dbURL := os.Getenv("ACCOUNT_DB_URL")
	if dbURL == "" {
		log.Fatal("ACCOUNT_DB_URL environment variable not set")
	}

	repo, err := account.NewPostgresRepository(dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer repo.Close()

	service := account.NewService(repo)
	    // Start the gRPC server in a goroutine
    go func() {
        log.Println("Account gRPC server starting on port 8080...")
        if err := account.ListenAndServeGRPC(service, 8080); err != nil {
            log.Fatalf("Failed to start gRPC server: %v", err)
        }
    }()

    // Wait for an interrupt signal to gracefully shut down
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    log.Println("Connected to DB. Server is running. Press Ctrl+C to stop.")
    <-quit
    log.Println("Shutting down account service...")
}
