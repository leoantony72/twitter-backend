package model

import (
	"time"
)

type User struct {
	Id           string    `json:"id" gorm:"primaryKey;type:text"`
	Username     string    `json:"username" gorm:"type:varchar(25);not null"`
	Email        string    `json:"email" gorm:"type:varchar(40);uniqueIndex;not null"`
	Password     string    `json:"password" gorm:"type:varchar(50);not null"`
	Salt         string    `gorm:"type:text"`
	Token        string    `gorm:"type:text"`
	Date_created time.Time `json:"data_created"`
}
