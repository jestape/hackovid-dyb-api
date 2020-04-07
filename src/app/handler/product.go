package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/jestape/hackovid-dyb-api/src/app/model"
)

func GetProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	product := model.Product{}
	params := mux.Vars(r)

	if err := db.Where("id = ?", params["id"]).Find(&product).Error; err != nil {
		respondError(w, http.StatusOK, "Product not found") 
	} else {
		respondJSON(w, http.StatusOK, product) 
	}
	
}

func GetProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	products := []model.Product{}

	db.Find(&products)
	respondJSON(w, http.StatusOK, products)
	
}

func CreateProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	var product model.Product = model.Product{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	
	if (!checkNotNil(product.Name, "name", w)) { 
		return 
	}

	if err := db.Save(&product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, product)

}
