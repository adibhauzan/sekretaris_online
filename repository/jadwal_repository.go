package repository

import (
	"github.com/adibhauzan/sekretaris_online_backend/models"
	"time"
)

type JadwalRepository interface {
	CreateJadwal(jadwal *models.Jadwal) error
	GetAllJadwal() ([]models.Jadwal, error)
	GetJadwalByID(id uint) (*models.Jadwal, error)
	UpdateJadwal(jadwal *models.Jadwal) error
	DeleteJadwal(id uint) error
	GetJadwalByDatetime(date time.Time) ([]models.Jadwal, error)
}
