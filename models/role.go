package models

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Permissions []Permission `gorm:"ForeignKey:PermissionRefer"`
}

type UserRole struct {
	gorm.Model
	Users []User `gorm:"ForeignKey:UserRefer"`
	Roles []Role `gorm:"ForeignKey:RoleRefer"`
}
