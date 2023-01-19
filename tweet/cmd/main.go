package main

import (
	"github.com/leoantony72/twitter-backend/tweet/database"
	"github.com/leoantony72/twitter-backend/tweet/internal/repositories"
)

func main() {
	db := database.StartPostgres()
	repositories.NewTweetRepo(db)
}
