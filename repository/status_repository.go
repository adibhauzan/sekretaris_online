package repository

import "github.com/adibhauzan/sekretaris_online_backend/models"

type StatusRepository interface {
    CreateStatus(status *models.Status) error
    GetAllStatus() ([]models.Status, error)
    GetStatusByID(id uint) (*models.Status, error)
    UpdateStatus(status *models.Status) error
    DeleteStatus(id uint) error
}
