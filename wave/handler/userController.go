package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wave/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []model.User{}
	db.Find(&users)
	respondJSON(w, http.StatusOK, users)
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	user := model.User{}

	json.NewDecoder(r.Body).Decode(&user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password Encryption  failed",
		}
		json.NewEncoder(w).Encode(err)
	}

	user.Password = string(pass)

	createdUser := db.Create(user)
	var errMessage = createdUser.Error

	if createdUser.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdUser)
}
