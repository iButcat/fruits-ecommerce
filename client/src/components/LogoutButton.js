import React, { useEffect, refetch } from 'react';
import { Nav, Button } from 'react-bootstrap';
import { Redirect } from 'react-router-dom';
 
function LogoutButton() {
    
    const handleOnClick = (event) => {
        event.preventDefault();
        localStorage.removeItem('token');
        localStorage.setItem('logged', JSON.stringify(false));
        console.log("LOGGED: ", localStorage.getItem('logged'));
        return <Redirect to="/home" />;
    };

    useEffect(() => {
        window.addEventListener("storage", () => {
            refetch();
        });
    });

    return (
        <Nav.Link onClick={(e) => handleOnClick(e)}>
             <Button variant="primary" type="submit">
                 Logout
             </Button>
        </Nav.Link>
    );
}

export default LogoutButton;