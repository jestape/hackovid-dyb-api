package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model

	PublicKey 		string  		`gorm:"primary_key" json:"public_key"`
	Name  			string     		`json:"name"`
	Email	 		string			`json:"email"`
	NIF				string			`json:"NIF"`

}

type Product struct {
	gorm.Model

	Id		 		int  			`gorm:"primary_key;AUTO_INCREMENT json:"product_id"`
	Name  			string     		`json:"name"`
	Price	 		int				`json:"email"`

}

type Tickets struct {

	Seller		User  			`gorm:"primary_key" json:"seller"` 
	Buyer 		User  			`gorm:"primary_key" json:"buyer"`
	Timestamp	time.Time		`gorm:"primary_key" json:"timestamp"`
	Products	[]Product		`json:"buyer"`
	TotalPrice	int				`json:"total_price"`
	Hash		string			`json:"integrity_proof"`

}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&SensorData{})	
	return db
}
