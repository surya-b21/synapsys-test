package service

import (
	"fmt"
	"reflect"
	"synapsis-test/app/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// SeedAll data
func SeedAll(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {
		seeds := DataSeeds()
		for i := range seeds {
			err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(seeds[i]).Error
			if nil != err {
				name := reflect.TypeOf(seeds[i]).String()
				errorMessage := err.Error()
				return fmt.Errorf("%s seeder fail with %s", name, errorMessage)
			}
		}
		return nil
	})
}

var (
	costumer        model.Costumer
	productCategory model.ProductCategory
	product         model.Product
)

// DataSeeds data to seed
func DataSeeds() []interface{} {
	return []interface{}{
		costumer.Seed(),
		productCategory.Seed(),
		product.Seed(),
	}
}
