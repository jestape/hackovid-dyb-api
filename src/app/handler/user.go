package handler

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/jestape/hackovid-dyb-api/src/app/model"
)

func GetUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	users := []model.User{}

	db.Find(&users)
	respondJSON(w, http.StatusOK, users)
	
}


func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, user)

}
