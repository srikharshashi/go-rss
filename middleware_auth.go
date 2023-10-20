package main

import (
	"fmt"
	"net/http"

	"github.com/srikharshashi/go-rss/internal/auth"
	"github.com/srikharshashi/go-rss/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User) 

func (cfg *ApiConfig) MiddleWareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, fmt.Sprintf("Auth Error %v", err), 403)
			return
		}

		db_user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, fmt.Sprintf("DB Error %v", err), 400)
			return
		}

		handler(w,r,db_user)
	}

}
