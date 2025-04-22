package repository

import (
	"gorm.io/gorm"
	"hospital-management-system/internal/models"
)

// PatientRepository handles database operations for patients.
type PatientRepository struct {
	DB *gorm.DB
}

// NewPatientRepository creates a new instance of PatientRepository.
func NewPatientRepository(db *gorm.DB) *PatientRepository {
	return &PatientRepository{DB: db}
}

// CreatePatient adds a new patient to the database.
func (r *PatientRepository) CreatePatient(patient *models.Patient) error {
	return r.DB.Create(patient).Error
}

// GetPatient retrieves a patient by ID.
func (r *PatientRepository) GetPatient(id uint) (*models.Patient, error) {
	var patient models.Patient
	if err := r.DB.First(&patient, id).Error; err != nil {
		return nil, err
	}
	return &patient, nil
}

// UpdatePatient updates an existing patient's details.
func (r *PatientRepository) UpdatePatient(patient *models.Patient) error {
	return r.DB.Save(patient).Error
}

// DeletePatient removes a patient from the database.
func (r *PatientRepository) DeletePatient(id uint) error {
	return r.DB.Delete(&models.Patient{}, id).Error
}

// ListPatients retrieves all patients from the database.
func (r *PatientRepository) ListPatients() ([]models.Patient, error) {
	var patients []models.Patient
	if err := r.DB.Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}