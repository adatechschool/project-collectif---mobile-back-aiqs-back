package model

import "gorm.io/gorm"

type User struct {
	//gorm.Model
	UserID   int `gorm:"primary_key;AUTO_INCREMENT"`
	Name     string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `json:"Password"`
	//Spots    Surf
}

func DBMigrateUser(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}
