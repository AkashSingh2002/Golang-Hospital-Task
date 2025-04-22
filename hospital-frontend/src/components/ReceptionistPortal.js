import React, { useState, useEffect } from 'react';
import './ReceptionistPortal.css';

function ReceptionistPortal() {
  const [patients, setPatients] = useState([]);
  const [newPatient, setNewPatient] = useState({
    firstName: '',
    lastName: '',
    age: '',
    gender: '',
    address: '',
    phone: '',
  });

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

  const handleAddPatient = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('http://localhost:8080/patients', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${localStorage.getItem('token')}`,
        },
        body: JSON.stringify(newPatient),
      });
      if (response.ok) {
        fetchPatients();
        setNewPatient({ firstName: '', lastName: '', age: '', gender: '', address: '', phone: '' });
      } else {
        alert('Failed to add patient');
      }
    } catch (error) {
      console.error('Error adding patient:', error);
    }
  };

  return (
    <div className="receptionist-portal">
      <h2>Receptionist Portal</h2>
      <form onSubmit={handleAddPatient}>
        <h3>Add New Patient</h3>
        <input
          type="text"
          placeholder="First Name"
          value={newPatient.firstName}
          onChange={(e) => setNewPatient({ ...newPatient, firstName: e.target.value })}
        />
        <input
          type="text"
          placeholder="Last Name"
          value={newPatient.lastName}
          onChange={(e) => setNewPatient({ ...newPatient, lastName: e.target.value })}
        />
        <input
          type="number"
          placeholder="Age"
          value={newPatient.age}
          onChange={(e) => setNewPatient({ ...newPatient, age: e.target.value })}
        />
        <input
          type="text"
          placeholder="Gender"
          value={newPatient.gender}
          onChange={(e) => setNewPatient({ ...newPatient, gender: e.target.value })}
        />
        <input
          type="text"
          placeholder="Address"
          value={newPatient.address}
          onChange={(e) => setNewPatient({ ...newPatient, address: e.target.value })}
        />
        <input
          type="text"
          placeholder="Phone"
          value={newPatient.phone}
          onChange={(e) => setNewPatient({ ...newPatient, phone: e.target.value })}
        />
        <button type="submit">Add Patient</button>
      </form>

      <h3>Patient List</h3>
      <ul>
        {patients.map((patient) => (
          <li key={patient.id}>{`${patient.firstName} ${patient.lastName}`}</li>
        ))}
      </ul>
    </div>
  );
}

export default ReceptionistPortal;