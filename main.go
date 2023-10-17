package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/srikharshashi/go-projects/go-rss/internal/database"
)

// this type database.Queries is exposed by SQLC generate code in internal
// has ro be added manually
type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT IS NOT FOUND IN ENV")
	}

	DBURL := os.Getenv("DBURL")
	if DBURL == "" {
		log.Fatal("DBURL IS NOT FOUND IN ENV")
	}


	db, err := sql.Open("mysql", DBURL)
	if err != nil {
		
		log.Fatal("DB Connection Failed", err)
	}

	queries := database.New(db)
	apiCfg := apiConfig{
		DB: queries,
	}

	fmt.Println("Starting server on PORT", PORT)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerError)
	v1Router.Get("/error", handlerError)
	v1Router.Post("/users", apiCfg.handlerCreateUser)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	log.Printf("Sever starting on prt %v", PORT)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error in Server")

	}

}
