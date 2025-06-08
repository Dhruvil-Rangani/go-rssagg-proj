package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/cors"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load();
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err);
	}

	port := os.Getenv("PORT");
	if port == "" {
		log.Fatal("PORT environment variable is not set");
	}

	router := chi.NewRouter();

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*","http://*"}, // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Content-Length","Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum age for preflight requests
	}))

	v1Router := chi.NewRouter();
	v1Router.Get("/healthz", handlerReadiness);
	v1Router.Get("/error", handlerErr);
	router.Mount("/v1", v1Router);


	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", port),
	}

	log.Printf("Server will run on port: %s\n", port);

	err = srv.ListenAndServe();
	if err != nil {
		log.Fatalf("Error starting server: %v", err);
	}

}