package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"blog.com/go/rest-ws/handlers"
	"blog.com/go/rest-ws/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWT_SECRET,
		Port:        PORT,
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BinderRoutes)
}

func BinderRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
}
