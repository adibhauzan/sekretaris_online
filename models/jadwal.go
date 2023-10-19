package models

import (
	"gorm.io/gorm"
	"time"
)

type Jadwal struct {
	gorm.Model
	Name       string    `json:"name"`
	Date       time.Time `json:"date"`
	StatusID   uint      `json:"status_id"`
	Status     Status    `json:"status"`
	ExpiryDate time.Time `gorm:"-"`
}
