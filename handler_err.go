package main

import "net/http"

//just a function to respond to errors(demo)

func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, "Something went wrong", 400)
}
