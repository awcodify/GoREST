package models

import (
	"github.com/jinzhu/gorm"
)

type Permission struct {
	gorm.Model
	Keyname     string `json:"keyName"`
	Description string `json:"description"`
}
