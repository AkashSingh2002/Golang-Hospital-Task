package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"hospital-management-system/internal/models"
	"hospital-management-system/internal/service"
	"github.com/gorilla/mux"
)

// PatientHandler handles HTTP requests for patient management.
type PatientHandler struct {
	Service *service.PatientService
}

// NewPatientHandler creates a new instance of PatientHandler.
func NewPatientHandler(service *service.PatientService) *PatientHandler {
	return &PatientHandler{Service: service}
}

// CreatePatient handles the creation of a new patient.
func (h *PatientHandler) CreatePatient(w http.ResponseWriter, r *http.Request) {
	var patient models.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.Service.CreatePatient(&patient); err != nil {
		http.Error(w, "Failed to create patient", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// GetPatient handles retrieving a patient by ID.
func (h *PatientHandler) GetPatient(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	patient, err := h.Service.GetPatient(uint(id))
	if err != nil {
		http.Error(w, "Patient not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(patient)
}

// UpdatePatient handles updating an existing patient's details.
func (h *PatientHandler) UpdatePatient(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var patient models.Patient
	if err := json.NewDecoder(r.Body).Decode(&patient); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	patient.ID = uint(id)
	if err := h.Service.UpdatePatient(&patient); err != nil {
		http.Error(w, "Failed to update patient", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeletePatient handles deleting a patient by ID.
func (h *PatientHandler) DeletePatient(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.Service.DeletePatient(uint(id)); err != nil {
		http.Error(w, "Failed to delete patient", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// ListPatients handles retrieving all patients.
func (h *PatientHandler) ListPatients(w http.ResponseWriter, r *http.Request) {
	patients, err := h.Service.ListPatients()
	if err != nil {
		http.Error(w, "Failed to retrieve patients", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(patients)
}