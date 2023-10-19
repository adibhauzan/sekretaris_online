package handlers

import (
	"time"

	"github.com/adibhauzan/sekretaris_online_backend/models"
	"gorm.io/gorm"
)

func AutoDeleteExpiredData(db *gorm.DB) {
	for {
		currentDate := time.Now()

		var expiredData []models.Jadwal
		db.Where("expiry_date < ?", currentDate).Find(&expiredData)

		for _, data := range expiredData {
			db.Delete(&data)
		}

		time.Sleep(24 * time.Hour)
	}
}
