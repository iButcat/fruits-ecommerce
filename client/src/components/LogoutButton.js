import React from 'react';
import { Link } from 'react-router-dom';
import { Button, Nav } from 'react-bootstrap';
 
function LogoutButton() {
    
    const handleOnClick = (event) => {
        event.preventDefault();
        localStorage.removeItem('token');
        localStorage.setItem('logged', JSON.stringify(false));
        console.log("LOGGED: ", localStorage.getItem('logged'));
    };

    return (
        <Nav.Link onClick={(e) => handleOnClick(e)} to="/home">
            Logout
        </Nav.Link>
    );
}

export default LogoutButton;