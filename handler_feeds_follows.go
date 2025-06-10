package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Dhruvil-Rangani/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedsFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type params struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	p := params{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request body: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedsFollows(r.Context(), database.CreateFeedsFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    p.FeedID,
	})
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to create feed follow: %v", err))
		return
	}

	jsonResponse(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedsFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollow, err := apiCfg.DB.GetFeedsFollows(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to get feed follow: %v", err))
		return
	}

	jsonResponse(w, http.StatusOK, databaseFeedsFollowsToFeedsFollows(feedFollow))
}

func (apiCfg *apiConfig) handlerDeleteFeedsFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID");
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid feed follow ID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedsFollows(r.Context(), database.DeleteFeedsFollowsParams{
		UserID: user.ID,
		ID: feedFollowID,
	})
	if err != nil {
		responseWithError(w, http.StatusBadRequest, fmt.Sprintf("Failed to delete feed follow: %v", err))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
