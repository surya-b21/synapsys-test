package model

import "strings"

// Product model
type Product struct {
	Base
	ProductCategory *ProductCategory `json:"product_category,omitempty" gorm:"foreignKey:ProductCategoryCode;references:Code"`
	ProductAPI
}

// ProductAPI detail
type ProductAPI struct {
	Name                *string `json:"name,omitempty" gorm:"type:varchar(100)"`
	ProductCategoryCode *string `json:"product_category_code,omitempty" gorm:"type:varchar(3)"`
}

// Seed data
func (s *Product) Seed() *[]Product {
	data := []Product{}
	items := []string{
		"Ayam|FNB",
		"Ikan|FNB",
		"Smarthphone|ELC",
		"Kamera|ELC",
		"Ban Motor|AUT",
		"Knalpot|AUT",
	}

	for i := range items {
		c := strings.Split(items[i], "|")

		name := c[0]
		code := c[1]

		data = append(data, Product{ProductAPI: ProductAPI{
			Name:                &name,
			ProductCategoryCode: &code,
		}})
	}

	return &data
}
