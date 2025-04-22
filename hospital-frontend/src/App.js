import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Login from './components/Login';
import ReceptionistPortal from './components/ReceptionistPortal';
import DoctorPortal from './components/DoctorPortal';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/receptionist" element={<ReceptionistPortal />} />
        <Route path="/doctor" element={<DoctorPortal />} />
      </Routes>
    </Router>
  );
}

export default App;
