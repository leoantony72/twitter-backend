package database

import (
	"log"

	"github.com/leoantony72/twitter-backend/timeline/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartPostgres() *gorm.DB {
	dbUrl := "postgres://pg:pass@database:5432/twitter"

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.User{})
	return db
}
