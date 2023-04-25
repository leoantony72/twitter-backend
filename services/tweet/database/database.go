package database

import (
	"log"

	"github.com/leoantony72/twitter-backend/tweet/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartPostgres() *gorm.DB {
	dbUrl := "postgres://pg:pass@database:5432/twitter"

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Tweets{})
	db.AutoMigrate(&model.Retweet{})
	db.AutoMigrate(&model.Like{})
	return db
}
