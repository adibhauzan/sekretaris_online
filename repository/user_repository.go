package repository

import "github.com/adibhauzan/sekretaris_online_backend/models"

type UserRepository interface {
    CreateUser(user *models.User) error
    GetUserByUsername(username string) (*models.User, error)
}

