package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// NewDatabase created new instance to database...
func NewDatabase(host string, port string, user string, pass string, base string) *gorm.DB {
	cfg := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, base, pass)
	db, err := gorm.Open("postgres", cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db
}

// Migrates create my tables
func Migrates(db *gorm.DB) {
	db.AutoMigrate(&Account{}, &OperationType{}, &Transaction{})

	db.Create(&OperationType{Description: CompraAVista})
	db.Create(&OperationType{Description: CompraParcelada})
	db.Create(&OperationType{Description: Saque})
	db.Create(&OperationType{Description: Pagamento})
}
