package main

import (
	"fmt"
	"net/http"

	"github.com/emresoysuren/rssagg/internal/auth"
	"github.com/emresoysuren/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintln("Auth error:", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintln("Couldn't get user:", err))
			return
		}

		handler(w, r, user)
	}
}
