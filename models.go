package main

import (
	"time"

	"github.com/srikharshashi/go-rss/internal/database"
)

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func databaseUsertoUser(dbUser database.User) User{
	return User{
		ID:dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
	}

}