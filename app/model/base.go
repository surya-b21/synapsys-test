package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base struct
type Base struct {
	ID        *uuid.UUID     `json:"id,omitempty" gorm:"primaryKey;unique;type:varchar(36);not null" format:"uuid"`
	CreatedAt *time.Time     `json:"created_at,omitempty" gorm:"type:timestamp" format:"date-time" swaggertype:"string"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty" gorm:"type:timestamp" format:"date-time" swaggertype:"string"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`
}

// BeforeCreate hooks
func (b *Base) BeforeCreate(tx *gorm.DB) error {
	if b.ID != nil {
		return nil
	}
	id, e := uuid.NewRandom()
	now := time.Now().UTC().Add(7 * time.Hour)
	b.ID = &id
	b.CreatedAt = &now
	b.UpdatedAt = &now
	return e
}

// BeforeUpdate hooks
func (b *Base) BeforeUpdate(tx *gorm.DB) error {
	if b.ID != nil {
		return nil
	}
	now := time.Now().UTC().Add(7 * time.Hour)
	b.UpdatedAt = &now
	return nil
}
