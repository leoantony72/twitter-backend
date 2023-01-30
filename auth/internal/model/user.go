package model

import "time"

type User struct {
	Id           string    `json:"id" gorm:"primaryKey;type:text"`
	Username     string    `json:"username" gorm:"type:varchar(25);uniqueIndex;not null"`
	Email        string    `json:"email" gorm:"type:varchar(40);uniqueIndex;not null"`
	Password     string    `json:"password,omitempty" gorm:"type:varchar(50);not null"`
	Followers    int       `json:"followers" gorm:"type:integer;default:0"`
	Following    int       `json:"following" gorm:"type:integer;default:0"`
	Salt         string    `json:"-" gorm:"type:text"`
	Token        string    `json:"-" gorm:"type:text"`
	Date_created time.Time `json:"data_created" gorm:"type:timestamp"`
}

type User_followers struct {
	Follower string `json:"follower" gorm:"primaryKey;type:varchar(55) REFERENCES users(username);not null"`
	Followee string `json:"followee" gorm:"primaryKey;type:varchar(55) REFERENCES users(username);not null"`
}
