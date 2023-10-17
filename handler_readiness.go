package main

import "net/http"

//just a function to respond to requests to check if the server is ready

func handlerReadiness(w http.ResponseWriter,r *http.Request) {
	respondWithJSON(w,200,struct{}{})
}