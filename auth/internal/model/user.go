package model

import "time"

type User struct {
	Id           string    `json:"id" gorm:"primaryKey"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Date_created time.Time `json:"data_created"`
}
