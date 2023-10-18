package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/srikharshashi/go-rss/internal/database"
)


func genAPIKEY() string {
	//8bytes = 64 bits
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Error generating random bytes:", err)
	}
	// Convert the random bytes to a hexadecimal string
	return hex.EncodeToString(randomBytes)

}

//here what we have done is we made this function a methods for the type apiConfig instead
//this can be called with an object of api config instead now
func (apiCfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter,r *http.Request) {
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
	api_key:= genAPIKEY()

	err= apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID:       id ,
		CreatedAt: time_now,
		UpdatedAt: time_now,
		Name:      params.Name,
		ApiKey: api_key,
	})
	res:=User{
		ID: id,
		CreatedAt: time_now,
		UpdatedAt: time_now,
		Name: params.Name,
		ApiKey: api_key,
	}
	if(err!=nil){
		respondWithError(w,fmt.Sprintf("POST-/user couldn't create a user %v",err),400)
		return 
	}

	respondWithJSON(w,201,res);

}


func (apiCfg *ApiConfig) handlerGetUser(w http.ResponseWriter,r *http.Request,db_user database.User) {
	 respondWithJSON(w,200,databaseUsertoUser(db_user))
}