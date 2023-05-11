package service

import "gorm.io/gorm"

// AutoMigrate function
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(ModelList...)
}

// ModelList var
var ModelList []interface{} = []interface{}{}
