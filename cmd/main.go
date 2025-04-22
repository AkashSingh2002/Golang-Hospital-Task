package main

import (
	"log"
	"net/http"

	"hospital-management-system/internal/handlers"
	"hospital-management-system/internal/models"
	"hospital-management-system/internal/repository"
	"hospital-management-system/internal/service"
	"hospital-management-system/pkg/config"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := gorm.Open(sqlite.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate models
	db.AutoMigrate(&models.Patient{})

	// Initialize repositories, services, and handlers
	patientRepo := repository.NewPatientRepository(db)
	patientService := service.NewPatientService(patientRepo)
	patientHandler := handlers.NewPatientHandler(patientService)

	// Set up router
	r := mux.NewRouter()
	r.HandleFunc("/patients", patientHandler.CreatePatient).Methods("POST")
	r.HandleFunc("/patients", patientHandler.ListPatients).Methods("GET")
	r.HandleFunc("/patients/{id}", patientHandler.GetPatient).Methods("GET")
	r.HandleFunc("/patients/{id}", patientHandler.UpdatePatient).Methods("PUT")
	r.HandleFunc("/patients/{id}", patientHandler.DeletePatient).Methods("DELETE")

	// Start server
	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}