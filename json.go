package main

import (
	"encoding/json"
	"log"
	"net/http"
)


func respondWithJSON(w http.ResponseWriter,code int,payload interface{}) {

	data,err:=json.Marshal(payload)
	if(err!=nil){
		log.Printf("JSON MiddleWare:Error Marshalling %v",payload);
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter,errorMessage string,code int){
	if(code >499){
		//client error
		log.Printf("ERROR:Client Side Error Code : %v Error: %v ",code,errorMessage);

	}

	//uses json reflect tags for the json library to marshall things properly
	type errorResponse struct{
		Error string `json:"error"`

	}

	respondWithJSON(w,code,errorResponse{
		Error: errorMessage,
	})


}