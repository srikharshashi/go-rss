package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	
	godotenv.Load()
	port:= os.Getenv("PORT")
	if(port=="") {
		log.Fatal("PORT IS NOT FOUND IN ENV")
	}
	
	fmt.Println("Starting server on PORT",port);

	router:=chi.NewRouter()
	srv:= &http.Server{
		Handler: router,
		Addr:":"+port,
	}

	router.Use(
		cors.Handler(
			cors.Options{
				
			}
		)
	)

	log.Printf("Sever starting on prt %v",port)
	err:=srv.ListenAndServe()
	if(err!=nil){
		log.Fatal("Error in Server");
	}
	
}