import React, { useState, useEffect } from 'react';
import './DoctorPortal.css';

function DoctorPortal() {
  const [patients, setPatients] = useState([]);

  useEffect(() => {
    fetchPatients();
  }, []);

  const fetchPatients = async () => {
    try {
      const response = await fetch('http://localhost:8080/patients', {
        headers: { Authorization: `Bearer ${localStorage.getItem('token')}` },
      });
      const data = await response.json();
      setPatients(data);
    } catch (error) {
      console.error('Error fetching patients:', error);
    }
  };

  const handleUpdatePatient = async (id, updatedData) => {
    try {
      const response = await fetch(`http://localhost:8080/patients/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${localStorage.getItem('token')}`,
        },
        body: JSON.stringify(updatedData),
      });
      if (response.ok) {
        fetchPatients();
      } else {
        alert('Failed to update patient');
      }
    } catch (error) {
      console.error('Error updating patient:', error);
    }
  };

  return (
    <div className="doctor-portal">
      <h2>Doctor Portal</h2>
      <h3>Patient List</h3>
      <ul>
        {patients.map((patient) => (
          <li key={patient.id}>
            {`${patient.firstName} ${patient.lastName}`} - {patient.age} years old
            <button
              onClick={() =>
                handleUpdatePatient(patient.id, {
                  ...patient,
                  age: patient.age + 1, // Example update: increment age
                })
              }
            >
              Update Age
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default DoctorPortal;