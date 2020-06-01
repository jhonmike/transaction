package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// CashPurchase constant representing a operation type
const CashPurchase string = "COMPRA A VISTA"

// InstallmentPurchase constant representing a operation type
const InstallmentPurchase string = "COMPRA PARCELADA"

// Withdraw constant representing a operation type
const Withdraw string = "SAQUE"

// Payment constant representing a operation type
const Payment string = "PAGAMENTO"

// OperationType model that represents the types of operations transacted in the account
type OperationType struct {
	ID           uint          `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	DeletedAt    *time.Time    `sql:"index" json:"deleted_at"`
	Description  string        `json:"description"`
	Transactions []Transaction `gorm:"foreignkey:OperationTypeID" json:"_"`
}

// OperationTypeResource interface
type OperationTypeResource interface {
	GetOperationTypeByID(operationType OperationType, ID interface{}) (OperationType, error)
}

type operationTypeResource struct {
	db *gorm.DB
}

// NewOperationTypeResource created new instance of the OperationTypes model manipulation resource
func NewOperationTypeResource(db *gorm.DB) OperationTypeResource {
	return operationTypeResource{db}
}

func (resource operationTypeResource) GetOperationTypeByID(operationType OperationType, ID interface{}) (OperationType, error) {
	if err := resource.db.First(&operationType, ID).Error; err != nil {
		return operationType, err
	}
	return operationType, nil
}
