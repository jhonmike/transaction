package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Account model representing the customer is back account
type Account struct {
	ID             uint          `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	DeletedAt      *time.Time    `sql:"index" json:"deleted_at"`
	DocumentNumber string        `json:"document_number"`
	Transactions   []Transaction `gorm:"foreignkey:AccountID" json:"_"`
}

// AccountResource interface
type AccountResource interface {
	CreateAccount(account Account) (Account, error)
	GetAccountByID(account Account, ID interface{}) (Account, error)
}

type accountResource struct {
	db *gorm.DB
}

// NewAccountResource created new instance of the Account model manipulation resource
func NewAccountResource(db *gorm.DB) AccountResource {
	return accountResource{db}
}

func (resource accountResource) CreateAccount(account Account) (Account, error) {
	if err := resource.db.Create(&account).Error; err != nil {
		return account, err
	}
	return account, nil
}

func (resource accountResource) GetAccountByID(account Account, ID interface{}) (Account, error) {
	if err := resource.db.First(&account, ID).Error; err != nil {
		return account, err
	}
	return account, nil
}
