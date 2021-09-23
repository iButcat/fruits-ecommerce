import React from 'react';
import { Link } from 'react-router-dom';
import { Button } from 'react-bootstrap';
 
function LogoutButton() {
    
    const handleOnClick = (event) => {
        event.preventDefault();
        localStorage.removeItem('token');
        localStorage.setItem('logged', JSON.stringify(false));
        console.log("LOGGED: ", localStorage.getItem('logged'));
    };

    return (
        <Button onClick={(e) => handleOnClick(e)}>
            <Link to="/home" style={{ textDecoration: 'none' }} >
            Logout
            </Link>
        </Button>
    );
}

export default LogoutButton;