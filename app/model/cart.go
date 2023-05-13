package model

import (
	"github.com/google/uuid"
)

// Cart models
type Cart struct {
	Base
	Product  *Product  `json:"product"`
	Costumer *Costumer `json:"costumer"`
	CartAPI
}

// CartAPI details
type CartAPI struct {
	ProductID  *uuid.UUID `json:"product_id,omitempty" gorm:"type:varchar(36);not null"`
	CostumerID *uuid.UUID `json:"costumer_id,omitempty" gorm:"type:varchar(36);not null"`
	IsCheckout *bool      `json:"is_checkout,omitempty" gorm:"type:bool"`
}
