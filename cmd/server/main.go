package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/prastamaha/envelope-encryption-example/internal/handler"
	"github.com/prastamaha/envelope-encryption-example/internal/repository"
	"github.com/prastamaha/envelope-encryption-example/internal/util/libdb"
	"github.com/prastamaha/envelope-encryption-example/internal/util/libkms"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	// Initialize Vault client
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	vault := libkms.NewVaultKMS(
		os.Getenv("VAULT_ADDR"),
		os.Getenv("VAULT_TOKEN"),
		tlsConfig,
	)

	client, err := vault.NewClient()
	if err != nil {
		log.Fatalf("Unable to initialize Vault client: %v", err)
	}
	log.Println("Vault client successfully initialized")

	// Initialize database
	postgres := libdb.NewPostgres(
		os.Getenv("DATABASE_HOSTNAME"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	)
	db := postgres.InitDB()
	log.Println("Database successfully initialized")

	// Initialize repository
	cryptoRepo := repository.NewCrypto(client, os.Getenv("VAULT_KEK_NAME"))
	userRepo := repository.NewUserRepository(db, cryptoRepo)
	log.Println("Repository successfully initialized")

	// Initialize http handler
	router := mux.NewRouter()
	httpHandler := handler.NewHandler(router, userRepo)
	httpHandler.RegisterRoutes()

	srv := &http.Server{
		Handler: router,
		Addr:    os.Getenv("SERVER_ADDR"),
	}

	log.Println("Server started on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
