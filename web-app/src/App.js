import React from 'react';
import { Switch, Route } from "react-router-dom";



//Views here
import Home from "./views/Home.view"
// import Reservation from "./views/Reservation.view"

//Css
import './css/App.css'
import './css/Menu/Navegation.css'
import './css/ReservationPage.css'


function App(props) {
  console.log(props)
  return (
    <Switch>
      <Route exact path="/" component={Home}/>
      
    </Switch>
   
  )
}

export default App
