package service

import (
	"hospital-management-system/internal/models"
	"hospital-management-system/internal/repository"
)

// PatientService provides business logic for patient management.
type PatientService struct {
	repo *repository.PatientRepository
}

// NewPatientService creates a new instance of PatientService.
func NewPatientService(repo *repository.PatientRepository) *PatientService {
	return &PatientService{repo: repo}
}

// CreatePatient handles the creation of a new patient.
func (s *PatientService) CreatePatient(patient *models.Patient) error {
	return s.repo.CreatePatient(patient)
}

// GetPatient retrieves a patient by ID.
func (s *PatientService) GetPatient(id uint) (*models.Patient, error) {
	return s.repo.GetPatient(id)
}

// UpdatePatient updates an existing patient's details.
func (s *PatientService) UpdatePatient(patient *models.Patient) error {
	return s.repo.UpdatePatient(patient)
}

// DeletePatient deletes a patient by ID.
func (s *PatientService) DeletePatient(id uint) error {
	return s.repo.DeletePatient(id)
}

// ListPatients retrieves all patients.
func (s *PatientService) ListPatients() ([]models.Patient, error) {
	return s.repo.ListPatients()
}