import React, { Component } from 'react';

class ReservationPage extends Component {
    constructor(props) {
        super(props);
        // this.state = {  };
    }
    render() {
        return (
            <div className="ReservationPage">
                <div class="container">
                    <h1>Reservation Page</h1>
                    <p>Please fill in this form to register for an event.</p>
                    <label for="first-name"><b>First Name</b></label>
            <input type="text" placeholder="First Name" name="first_name" required/>
            <br/>
            <label for="last-name"><b>Last Name</b></label>
            <input type="text" placeholder="Last Name" name="second_name" required/>
            <br/>
            <label for="username"><b>Username</b></label>
            <input type="text" placeholder="Enter Username" name="email" required/>
            <br/>
            <label for="email"><b>Email</b></label>
            <input type="text" placeholder="Enter Email" name="email" required/>
            </div>

            <br/>
 
            <label for="tickets" required> Select Tickets </label>
            <select name="tickets" id="tickets">
                <option value=""> </option>
                <option value="0">0</option>
                <option value="1">1</option>
            </select>
            <br/><br/>

            <label for="events" required>Select Event</label>
            <select name="events" id="events" required>
            <option value=""> </option>
                <option value="event1">Lehman College Hackathon</option>
                <option value="event2">Codelabs Summit</option>
                <option value="event3">Queens College Career Fair</option>
                <option value="event4">Island Trading</option>
                <option value="event5">Women in Tech Summer Program</option>
                <option value="event6">Modeling for Good </option>
                <option value="event6">Other</option>
            </select>
            <br/><br/>

            <label> If other </label><input type="text" placeholder="Enter Event Name"/>
            <br/><br/>
            {/* <!-- create submit botton --> */}
            <button type="button"> Make a Reservation</button>
            
    
                <div>
                    <h3>Winners</h3>
                <button type="button"> Run  Lottery</button>
                </div>
            </div>
            
        )
    }
}

export default ReservationPage;