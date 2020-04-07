package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/jestape/hackovid-dyb-api/src/app/model"
)

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	user := model.User{}
	params := mux.Vars(r)

	if err := db.Where(&model.User{PublicKey: params["pk"]}).Find(&user).Error; err != nil {
		respondError(w, http.StatusOK, "User not found") 
	} else {
		respondJSON(w, http.StatusOK, user) 
	}
	
}

func GetUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	users := []model.User{}

	db.Find(&users)
	respondJSON(w, http.StatusOK, users)
	
}

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	var user model.User = model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	
	if (!checkNotNil(user.PublicKey, "public_key", w) || !checkNotNil(user.Name, "user_name", w) ||
			!checkNotNil(user.Email, "email", w) || !checkNotNil(user.NIF, "nif", w)) { 
		return 
	}

	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, user)

}
