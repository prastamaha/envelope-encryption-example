package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/prastamaha/envelope-encryption-example/internal/handler"
	"github.com/prastamaha/envelope-encryption-example/internal/repository"
	"github.com/prastamaha/envelope-encryption-example/internal/util/libdb"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
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
	userRepo := repository.NewUserRepository(db)
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
