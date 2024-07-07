package models

import (
	"time"
)

type (
	Comments struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		Deskripsi string    `gorm:"not null" json:"deskripsi"`
		CreatedAt time.Time `gorm:"not null" json:"created_at"`
		UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
		UserID    uint      `gorm:"not null" json:"user_id"`
		ReviewID  uint      `gorm:"not null" json:"review_id"`
	}
)
