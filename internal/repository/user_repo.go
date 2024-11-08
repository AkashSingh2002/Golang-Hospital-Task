package repository

import (
	"database/sql"
	"hospital-management-system/internal/models"
)

func CreateUser(db *sql.DB, user *models.User) error {
	_, err := db.Exec("INSERT INTO users (username, password, role) VALUES ($1, $2, $3)", user.Username, user.Password, user.Role)
	return err
}

func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
	var user models.User
	err := db.QueryRow("SELECT id, username, password, role FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	return &user, err
}