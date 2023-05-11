package model

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Costumer models
type Costumer struct {
	Base
	CostumerAPI
}

// CostumerAPI data detail
type CostumerAPI struct {
	Name     *string `json:"name,omitempty" gorm:"type:varchar(100);not null"`
	Email    *string `json:"email,omitempty" gorm:"type:varchar(100);not null"`
	Password *string `json:"-" gorm:"type:varchar(1000);not null"`
}

// Seed data
func (s *Costumer) Seed() *[]Costumer {
	data := []Costumer{}
	items := []string{
		"admin|admin@email.com|admin12345",
	}

	for i := range items {
		c := strings.Split(items[i], "|")

		name := c[0]
		email := c[1]
		genpassword, _ := bcrypt.GenerateFromPassword([]byte(c[2]), 14)
		password := string(genpassword)

		data = append(data, Costumer{
			CostumerAPI: CostumerAPI{
				Name:     &name,
				Email:    &email,
				Password: &password,
			},
		})
	}

	return &data
}
