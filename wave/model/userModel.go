package model

import "gorm.io/gorm"

type User struct {
	//gorm.Model
	Name     string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `json:"Password"`
}

func DBMigrateUser(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}
