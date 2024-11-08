package migrations

import (
	"database/sql"
	"fmt"
	"log"
	"hospital-management-system/internal/models"
	"hospital-management-system/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate(db *sql.DB) {
	grmDb, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.LoadConfig().DBHost, config.LoadConfig().DBPort, config.LoadConfig().DBUser, config.LoadConfig().DBPassword, config.LoadConfig().DBName),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	err = grmDb.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Faileure : %v", err)
	}
	log.Println("Database migration completed successfully...")
}