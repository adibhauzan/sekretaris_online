package repositoryimpl

import (
	"time"

	"github.com/adibhauzan/sekretaris_online_backend/models"
	"gorm.io/gorm"
)

type JadwalRepositoryImpl struct {
	DB *gorm.DB
}

func NewJadwalRepository(db *gorm.DB) *JadwalRepositoryImpl {
	return &JadwalRepositoryImpl{DB: db}
}

func (r *JadwalRepositoryImpl) CreateJadwal(jadwal *models.Jadwal) error {
	return r.DB.Create(&jadwal).Error
}

func (r *JadwalRepositoryImpl) GetAllJadwal() ([]models.Jadwal, error) {
	var jadwals []models.Jadwal
	return jadwals, r.DB.Preload("Status").Find(&jadwals).Error
}

func (r *JadwalRepositoryImpl) GetJadwalByID(id uint) (*models.Jadwal, error) {
	var jadwal models.Jadwal
	return &jadwal, r.DB.Preload("Status").First(&jadwal, id).Error
}

func (r *JadwalRepositoryImpl) UpdateJadwal(jadwal *models.Jadwal) error {
	return r.DB.Save(&jadwal).Error
}

func (r *JadwalRepositoryImpl) DeleteJadwal(id uint) error {
	return r.DB.Delete(&models.Jadwal{}, id).Error
}

func (r *JadwalRepositoryImpl) GetJadwalByDatetime(datetime time.Time) ([]models.Jadwal, error) {
    var jadwals []models.Jadwal
    return jadwals, r.DB.Preload("Status").Where("date = ?", datetime).Find(&jadwals).Error
}