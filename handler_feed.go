package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Dhruvil-Rangani/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type params struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	p := params{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		Name: p.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url: p.URL,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to create user: %v", err))
		return
	}
	
	jsonResponse(w, http.StatusCreated, databaseUserToUser(feed))
}