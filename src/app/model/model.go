package model

import (
	//"time"

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

/*type Product struct {
	gorm.Model

	ProductId		uint  		`gorm:"primary_key;auto_increment:true"  json:"product_id"`
	Name  			string     	`json:"name"`
	Price	 		uint		`json:"email"`

}*/

/*type Ticket struct {

	Seller		User  			`gorm:"primary_key;auto_increment:false" json:"seller"` 
	Buyer 		User  			`gorm:"primary_key" json:"buyer"`
	Timestamp	time.Time		`gorm:"primary_key" json:"timestamp"`
	Products	[]Product		`json:"buyer"`
	TotalPrice	uint			`json:"total_price"`
	Hash		string			`json:"integrity_proof"`

}*/

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})	
	//db.AutoMigrate(&Product{})	
	//db.AutoMigrate(&Ticket{})
	return db
}
