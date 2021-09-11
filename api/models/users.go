package models

import "gorm.io/gorm"

// users list
type Users struct {
	Users []User `json:"users"`
}

// represent our user profile
type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}
