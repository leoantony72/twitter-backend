package model

import (
	"time"
)

type User struct {
	Id           string    `json:"id" gorm:"primaryKey;type:text"`
	Username     string    `json:"username" gorm:"type:varchar(25);uniqueIndex;not null"`
	Email        string    `json:"email" gorm:"type:varchar(40);uniqueIndex;not null"`
	Password     string    `json:"password" gorm:"type:varchar(50);not null"`
	Followers    int       `json:"followers" gorm:"type:integer;default:0"`
	Following    int       `json:"following" gorm:"type:integer;default:0"`
	Salt         string    `gorm:"type:text"`
	Token        string    `gorm:"type:text"`
	Date_created time.Time `json:"data_created"`
}

type User_followers struct {
	Follower string `json:"follower" gorm:"primaryKey;type:varchar(55) REFERENCES users(id);not null"`
	Followee string `json:"followee" gorm:"primaryKey;type:varchar(55) REFERENCES users(id);not null"`
}
