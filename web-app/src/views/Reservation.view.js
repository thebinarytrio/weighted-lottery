//cretae a react component
import React, { Component } from 'react';

//Import that component from components
import ReservationPage from "../components/ReservationPage"

//call the view 

class Reservation extends Component {
    render() {
        return (
            <div className="Reservation">
                <ReservationPage/>
            </div>
        );
    }
}

export default Reservation;
