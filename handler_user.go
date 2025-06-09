package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Dhruvil-Rangani/rssagg/internal/auth"
	"github.com/Dhruvil-Rangani/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Name string `json:"name"`
	}

	p := params{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		Name: p.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to create user: %v", err))
		return
	}
	
	jsonResponse(w, http.StatusCreated, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		responseWithError(w, http.StatusUnauthorized, fmt.Sprintf("Authentication failed: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		responseWithError(w, http.StatusNotFound, fmt.Sprintf("User not found: %v", err))
		return
	}
	jsonResponse(w, http.StatusOK, databaseUserToUser(user))
}