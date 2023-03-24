package models

import "pajaro.com/curso-go/db"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email	string `json:"email"`
}

type Users = []User

func MigrarUser() {
	db.Database.AutoMigrate(&User{})
}