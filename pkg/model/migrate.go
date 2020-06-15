package model

import (
	"github.com/jinzhu/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&User{}, &PublicKey{})
}
