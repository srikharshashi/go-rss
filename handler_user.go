package main

import (
	"encoding/json"
	"fmt"
	
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/srikharshashi/go-projects/go-rss/internal/database"
)

//here what we have done is we made this function a methods for the type apiConfig instead
//this can be called with an object of api config instead now
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter,r *http.Request) {
	type parameters struct{
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body);
	params:=parameters{}
	err:=decoder.Decode(&params)
	if err!=nil{
		respondWithError(w,fmt.Sprintf("Error parsing JSON",err),400)
		return
	}

	id:=uuid.New().String()
	time_now:=time.Now().UTC()
	err= apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID:       id ,
		CreatedAt: time_now,
		UpdatedAt: time_now,
		Name:      params.Name,
	})
	res:=User{
		ID: id,
		CreatedAt: time_now,
		UpdatedAt: time_now,
		Name: params.Name,
	}
	if(err!=nil){
		respondWithError(w,fmt.Sprintf("POST-/user couldn't create a user %v",err),400)
		return 
	}

	respondWithJSON(w,200,res);

}