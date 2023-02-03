package model

import (
	"time"
)

type User struct {
	Id           string    `json:"id" gorm:"primaryKey;type:text" redis:"id"`
	Username     string    `json:"username" gorm:"type:varchar(25);uniqueIndex;not null" redis:"username"`
	Email        string    `json:"email" gorm:"type:varchar(40);uniqueIndex;not null" redis:"email"`
	Password     string    `json:"password,omitempty" gorm:"type:varchar(50);not null" redis:"-"`
	Followers    int       `json:"followers" gorm:"type:integer;default:0" redis:"followers"`
	Following    int       `json:"following" gorm:"type:integer;default:0" redis:"following"`
	Salt         string    `json:"-" gorm:"type:text" redis:"-"`
	Token        string    `json:"-" gorm:"type:text" redis:"-"`
	Date_created time.Time `json:"data_created" gorm:"type:timestamp"`
	Encoded_Date string		`json:"-" redis:"date_created"`
}

type User_followers struct {
	Follower string `json:"follower" gorm:"primaryKey;type:varchar(55) REFERENCES users(username);not null"`
	Followee string `json:"followee" gorm:"primaryKey;type:varchar(55) REFERENCES users(username);not null"`
}
