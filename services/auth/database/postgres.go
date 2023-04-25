package database

import (
	"log"

	"github.com/leoantony72/twitter-backend/auth/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartPostgres() *gorm.DB {
	dbURL := "postgres://pg:pass@database:5432/twitter"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.User_followers{})
	return db
}
