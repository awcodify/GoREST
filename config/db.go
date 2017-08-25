package config

import (
	"github.com/jinzhu/gorm"
)

func Connect() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("mysql", "{username}:{password}@/{database}?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
