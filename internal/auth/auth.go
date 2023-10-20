package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string,error){

	val:= headers.Get("Authorization")
	if(val==""){
		return "",errors.New("authorization header not found")
	}

	vals:= strings.Split(val," ")
	if(len(vals)!=2){
		return "",errors.New("malformed auth header")
	}

	if vals[0] != "APIKEY"{
		return "",errors.New("malformed auth key (part1)")
	}

	return vals[1],nil

}