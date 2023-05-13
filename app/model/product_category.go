package model

import "strings"

// ProductCategory model
type ProductCategory struct {
	Base
	ProductCategoryAPI
}

// ProductCategoryAPI detail
type ProductCategoryAPI struct {
	Code     *string `json:"code,omitempty" gorm:"type:varchar(3);not null;uniqueIndex"`
	Category *string `json:"category,omitempty" gorm:"type:varchar(36);not null"`
}

// Seed data
func (s *ProductCategory) Seed() *[]ProductCategory {
	data := []ProductCategory{}
	items := []string{
		"FNB|Food and Beverage",
		"ELC|Electronic",
		"AUT|Automotive",
	}

	for i := range items {
		c := strings.Split(items[i], "|")

		code := c[0]
		category := c[1]

		data = append(data, ProductCategory{
			ProductCategoryAPI: ProductCategoryAPI{
				Code:     &code,
				Category: &category,
			},
		})
	}

	return &data
}
