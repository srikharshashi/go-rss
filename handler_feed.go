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
		respondWithError(w, fmt.Sprintf("Error parsing JSON %v", err), 400)
		return
	}

	id := uuid.New().String()
	time_now := time.Now().UTC()

	err = apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        id,
		CreatedAt: time_now,
		UpdatedAt: time_now,
		Name:      params.Name,
		Url: sql.NullString{
			String: params.Url,
			Valid:  true,
		},
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

func (apiconfig *ApiConfig) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, db_user database.User) {
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
		return
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

func (apiCfg *ApiConfig) handlerGetUseFollowFeed(w http.ResponseWriter, r *http.Request, db_user database.User) {

	db_feed_follows, err := apiCfg.DB.GetFeedFollowsByUser(r.Context(), db_user.ID)
	if err != nil {
		respondWithError(w, fmt.Sprintf("Error in Fetching Feeds user follows %v", err), 400)
		return
	}
	if len(db_feed_follows) == 0 {
		respondWithError(w, fmt.Sprintf("User doesn't follow any Feeds %v", err), 400)
		return

	}

	feeds := []Feed{}

	for _, feed := range db_feed_follows {
		db_feed, err := apiCfg.DB.GetFeedByID(r.Context(), feed.ID)
		if err != nil {
			respondWithError(w, fmt.Sprintf("Error in Fetching Feeds user follows %v", err), 400)
			return
		}

		feeds = append(feeds, Feed{
			ID:        db_feed.ID,
			CreatedAt: db_feed.CreatedAt,
			Name:      db_feed.Name,
			UpdatedAt: db_feed.UpdatedAt,
			Url:       db_feed.Url.String,
			User_id:   feed.UserID,
		})

	}

	respondWithJSON(w, 200, feeds)

}
