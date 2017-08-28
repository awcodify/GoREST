package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"username" validate:"min=3,max=40,regexp=^[a-zA-Z0-9]*$"`
	Password  string `json:"password" validate:"min=8"`
	Email     string `json:"email" validate:"nonzero"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

type TransformedUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
