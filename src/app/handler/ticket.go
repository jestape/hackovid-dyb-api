package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/jestape/hackovid-dyb-api/src/app/model"
)

func GetTicket(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	ticket := model.Ticket{}
	params := mux.Vars(r)
	db.Preload("TicketProducts.Product").Preload("Seller.User").Where("id = ?", params["id"]).Find(&ticket); 

	if ticket.SellerID == "" {
		respondError(w, http.StatusInternalServerError, "Ticket not found")
	} else {
		respondJSON(w, http.StatusOK, ticket) 
	}

}

func GetTicketTxHash(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	ticket := model.Ticket{}
	params := mux.Vars(r)
	db.Preload("TicketProducts.Product").Preload("Seller.User").Where("tx_hash = ?", params["tx_hash"]).Find(&ticket); 

	if ticket.SellerID == "" {
		respondError(w, http.StatusInternalServerError, "Ticket not found")
	} else {
		respondJSON(w, http.StatusOK, ticket) 
	}

}

func GetTickets(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	tickets := []model.Ticket{}

	db.Preload("TicketProducts.Product").Preload("Seller.User").Find(&tickets)
	respondJSON(w, http.StatusOK, tickets)
	
}

func GetTicketsUserB(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	buyer := r.URL.Query().Get("buyer")
	tickets := []model.Ticket{}
	db.Preload("TicketProducts.Product").Preload("Seller.User").Where("buyer_id = ?", buyer).Find(&tickets)
	respondJSON(w, http.StatusOK, tickets)

}

func GetTicketsUserS(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	seller := r.URL.Query().Get("seller")
	tickets := []model.Ticket{}
	db.Preload("TicketProducts.Product").Preload("Seller.User").Where("seller_id = ? AND tx_hash <> ''", seller).Find(&tickets)
	respondJSON(w, http.StatusOK, tickets)

}

func CreateTicket(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	var ticket model.Ticket = model.Ticket{}
	
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ticket); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&ticket).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	for _, row := range ticket.TicketProducts {
		ticket_product := model.TicketProduct{
			TicketID: ticket.ID,
			ProductID: row.ProductID,
			Amount: row.Amount,
		};
		if err := db.Save(&ticket_product).Error; err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	respondJSON(w, http.StatusCreated, ticket)

}

func UpdateTicket(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	var ticket model.Ticket = model.Ticket{}
	params := mux.Vars(r)

	db.Where("id = ?", params["id"]).Find(&ticket); 

	if (ticket.Hash != "") { 
		respondError(w, http.StatusBadRequest, "Ticket already processing or finalized")
		return 
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ticket); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if (!checkNotNil(ticket.Hash, "hash", w) || !checkNotNil(ticket.TxHash, "tx_hash", w) || !checkNotNil(ticket.BuyerID, "buyer", w)) { 
		return 
	}

	if err := db.Save(&ticket).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, ticket)

}
