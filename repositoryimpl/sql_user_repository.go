package repositoryimpl

import (
    "github.com/adibhauzan/sekretaris_online_backend/models"
    "gorm.io/gorm"
)

type UserRepositoryImpl struct {
    DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
    return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) CreateUser(user *models.User) error {
    return r.DB.Create(&user).Error
}

func (r *UserRepositoryImpl) GetUserByUsername(username string) (*models.User, error) {
    var user models.User
    return &user, r.DB.Where("username = ?", username).First(&user).Error
}

