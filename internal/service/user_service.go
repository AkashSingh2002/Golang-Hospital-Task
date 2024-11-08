package service

import (
	"database/sql"

	"hospital-management-system/internal/models"
	"hospital-management-system/internal/repository"
)

func CreateUser(db *sql.DB, user *models.User) error {
	return repository.CreateUser(db, user)
}

func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
	return repository.GetUserByUsername(db, username)
}