package main

import (
	"fmt"
	"net/http"

	"github.com/Dhruvil-Rangani/rssagg/internal/auth"
	"github.com/Dhruvil-Rangani/rssagg/internal/database"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg *apiConfig) authMiddleware(next authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		next(w, r, user)
	}
}	