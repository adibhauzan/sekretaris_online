package repositoryimpl

import (
    "github.com/adibhauzan/sekretaris_online_backend/models"
    "gorm.io/gorm"
)

type StatusRepositoryImpl struct {
    DB *gorm.DB
}

func NewStatusRepository(db *gorm.DB) *StatusRepositoryImpl {
    return &StatusRepositoryImpl{DB: db}
}

func (r *StatusRepositoryImpl) CreateStatus(status *models.Status) error {
    return r.DB.Create(&status).Error
}

func (r *StatusRepositoryImpl) GetAllStatus() ([]models.Status, error) {
    var statuses []models.Status
    return statuses, r.DB.Find(&statuses).Error
}

func (r *StatusRepositoryImpl) GetStatusByID(id uint) (*models.Status, error) {
    var status models.Status
    return &status, r.DB.First(&status, id).Error
}

func (r *StatusRepositoryImpl) UpdateStatus(status *models.Status) error {
    return r.DB.Save(&status).Error
}

func (r *StatusRepositoryImpl) DeleteStatus(id uint) error {
    return r.DB.Delete(&models.Status{}, id).Error
}
