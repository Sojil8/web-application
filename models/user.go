package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Name     string `gorm:"size(100)"`
	Email    string `gorm:"unique"`
	Password string `gorm:"not null"`
	Status   string `gorm:"type:varchar(50) ; check:status IN ('Active','Blocked')"`
}
