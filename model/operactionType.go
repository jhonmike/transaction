package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

const CompraAVista string = "COMPRA A VISTA"
const CompraParcelada string = "COMPRA PARCELADA"
const Saque string = "SAQUE"
const Pagamento string = "PAGAMENTO"

// OperationType model that represents the types of operations transacted in the account
type OperationType struct {
	ID           uint          `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	DeletedAt    *time.Time    `sql:"index" json:"deleted_at"`
	Description  string        `json:"description"`
	Transactions []Transaction `gorm:"foreignkey:OperationTypeID" json:"_"`
}

// OperationTypeResource ...
type OperationTypeResource interface {
	GetOperationTypeByID(operationType OperationType, ID interface{}) (OperationType, error)
}

type operationTypeResource struct {
	db *gorm.DB
}

// NewOperationTypeResource ...
func NewOperationTypeResource(db *gorm.DB) OperationTypeResource {
	return operationTypeResource{db}
}

func (resource operationTypeResource) GetOperationTypeByID(operationType OperationType, ID interface{}) (OperationType, error) {
	if err := resource.db.First(&operationType, ID).Error; err != nil {
		return operationType, err
	}
	return operationType, nil
}
