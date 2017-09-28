package database

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/whatdacode/GoREST/config"
	"github.com/whatdacode/GoREST/models"

	"gopkg.in/gormigrate.v1"
)

// Migrations is used for creater the init table for our database
func Migrations() *gormigrate.Gormigrate {
	db := config.Connect()
	db.LogMode(true)

	migrations := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: time.Now().Format("20060102") + "-users",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.User{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("users").Error
			},
		},
		{
			ID: time.Now().Format("20060102") + "-roles",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Role{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("roles").Error
			},
		},
		{
			ID: time.Now().Format("20060102") + "-permissions",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Permission{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("permissions").Error
			},
		}, {
			ID: time.Now().Format("20060102") + "-user_roles",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.UserRole{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("user_roles").Error
			},
		},
	})
	return migrations
}

// Migrate is used for migrating our database
func Migrate() {
	if err := Migrations().Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Printf("Migration did run successfully")
}

// Rollback is used for rollback our last migrations
func Rollback() {
	if err := Migrations().RollbackLast(); err != nil {
		log.Fatalf("Could not rollback: %v", err)
	}

	log.Printf("Rollback did run successfully")
}
