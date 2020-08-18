import React, {Component} from 'react'

class ReservationTable extends Component{

    constructor(props){
        super(props)
    }
    
    render(){
        return (
            <div className='ReservationTable'>
                <h1>List of Reservations</h1>
                <h2>Reservation Table</h2>
        <table class="events" border="2">
           
            <tr>
                <th>UserID</th>
                <th>Username</th>
                <th>Reservation Name </th>
            </tr>
            <tr>
                <td>1</td>
                <td>Juan23</td>
                <td>Lehman College Hackathon</td>
            </tr>
            <tr>
                <td>2</td>
                <td>Yasirisortiz</td>
                <td>Codelabs Summit </td>
            </tr>
            <tr>
                <td>3</td>
                <td>DallasJ</td>
                <td>Queens College Career Fair</td>
            </tr>
            <tr>
                <td>4</td>
                <td>Jhonsmith</td>
                <td>Island Trading</td>
            </tr>
            <tr>
                <td>5</td>
                <td>Mery90</td>
                <td>Women in Tech Summer Program </td>
            </tr>
        </table>

            </div>
        
        )
    }
}


export default ReservationTable