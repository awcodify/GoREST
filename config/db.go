package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Config struct {
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
}

func Connect() *gorm.DB {
	var configfile = "./config/db.conf"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open("mysql", config.DB_USER+":"+config.DB_PASSWORD+"@/"+config.DB_NAME+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	return db
}
