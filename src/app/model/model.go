package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {

	PublicKey 		string  			`gorm:"primary_key" json:"public_key"`
	Name  			string     			`gorm:"not null" json:"name"`
	Email	 		string				`gorm:"not null" json:"email"`
	NIF				string				`gorm:"not null" json:"nif"`
	Role			string				`gorm:"not null"`

}

type Product struct {
	
	ID				uint				`gorm:"primary_key" json:"product_id"`
	Name  			string     			`gorm:"not null" json:"name"`
	Price	 		uint				`gorm:"default:0" json:"price"`	

}

type Ticket struct {

	ID				uint				`json:"id"`

	Hash			string				`json:"hash"`
	TxHash			string				`json:"tx_hash"`
	Timestamp		time.Time			`json:"timestamp"`
	Seller      	*User 				`gorm:"foreignkey:SellerID; PRELOAD:true"`
	SellerID		string  			`json:"seller"`
	BuyerID			string  			`json:"buyer"` 
	
	TicketProducts	[]*TicketProduct    `gorm:"ForeignKey:TicketID" json:"products"`

}

type TicketProduct struct {
	
	TicketID		uint			 `gorm:"primary_key" json:"ticket_hash"`
	Product			*Product		
	ProductID		uint			 `gorm:"primary_key; auto_increment:false" json:"product_id"`
	
	Amount			uint			 `json:"amount"`

}

func (*TicketProduct) TableName() string {
    return "ticket_product"
}


// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})	
	db.AutoMigrate(&Product{})	
	db.AutoMigrate(&Ticket{})
	db.AutoMigrate(&TicketProduct{})
	return db
}
