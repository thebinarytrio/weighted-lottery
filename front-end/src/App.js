import React from 'react';
import logo from './logo.svg';
import './App.css';
import ReservationTable from './components/ReservationTable';
import ReservationPage from './components/ReservationPage';

function App() {
  return (
    <div className="App">
     <ReservationTable/>
     <ReservationPage/>
     
    </div>
  );
}

export default App;
