package model

import (
	//"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model

	PublicKey 		string  		`gorm:"unique" json:"public_key"`
	Name  			string     		`gorm:"not null" json:"name"`
	Email	 		string			`gorm:"not null" json:"email"`
	NIF				string			`gorm:"not null" json:"nif"`
	Role			string			`gorm:"default:'receiver'"`
	
}

/*type Product struct {
	gorm.Model

	ID				uint		
	Name  			string     	`json:"name"`
	Price	 		uint		`json:"price"`

}*/

/*type Ticket struct {
	gorm.Model

	Seller		string  		`gorm:"foreignkey:PublicKey" json:"seller"` 
	Buyer 		string  		`gorm:"foreignkey:PublicKey" json:"buyer"`
	Timestamp	time.Time		`json:"timestamp"`
	Products	[]Product		`gorm:"many2many:product_tickets" json:"products"`
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
