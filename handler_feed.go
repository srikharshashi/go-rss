package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/google/uuid"
	"github.com/srikharshashi/go-rss/internal/database"
)

// here what we have done is we made this function a methods for the type apiConfig instead
// this can be called with an object of api config instead now
func (apiCfg *ApiConfig) HandlerCreateFeed(w http.ResponseWriter, r *http.Request, db_user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Text string `json:"text"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, fmt.Sprintf("Error parsing JSON", err), 400)
		return
	}

	id := uuid.New().String()
	time_now := time.Now().UTC()

	err = apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        id,
		CreatedAt: time_now,
		UpdatedAt: time_now,
		Name:      params.Name,
		Url:       params.Url,
		UserID: sql.NullString{
			String: db_user.ID,
			Valid:  true,
		},
	})

	if err != nil {
		respondWithError(w, fmt.Sprintf("POST-/user couldn't create a user %v", err), 400)
		return
	}

	res := Feed{
		ID:        id,
		CreatedAt: time_now,
		UpdatedAt: time_now,
		Name:      params.Name,
		Url:       params.Url,
		User_id:   db_user.ID,
	}

	respondWithJSON(w, 201, res)

}

func (apicfg *ApiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apicfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, fmt.Sprintf("Error in Fetching feeds %v", err), 400)
	}
	respondWithJSON(w, 200, databaseFeedstoFeeds(feeds))

}

func (apiconfig *ApiConfig) HandlerFollowFeed(w http.ResponseWriter, r *http.Request, db_user database.User) {
	type parameters struct {
		FeedId string `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, fmt.Sprintf("Error parsing JSON %v", err), 400)
		return
	}

	id := uuid.New().String()
	time_now := time.Now().UTC()

	err = apiconfig.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        id,
		CreatedAt: time_now,
		UpdatedAt: time_now,
		UserID:    db_user.ID,
		FeedID:    params.FeedId,
	})

	if err != nil {
		respondWithError(w, fmt.Sprintf("Error in Adding Feed-Follow to database %v", err), 400)
	}

	res := FeedFollows{
		ID:        id,
		CreatedAt: time_now,
		UpdatedAt: time_now,
		UserID:    db_user.ID,
		FeedID:    params.FeedId,
	}

	respondWithJSON(w, 201, res)
}
