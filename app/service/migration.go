package service

import (
	"synapsis-test/app/model"

	"gorm.io/gorm"
)

// AutoMigrate function
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(ModelList...)
}

// ModelList var
var ModelList []interface{} = []interface{}{
	&model.Costumer{},
	&model.ProductCategory{},
	&model.Product{},
	&model.Cart{},
}
