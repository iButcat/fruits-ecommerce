import React from 'react';
import { Nav } from 'react-bootstrap';

function LoginButton() {
    return (
        <button className="LoginButton">
            <Nav.Link to="/login">Login</Nav.Link>
        </button>
    );
}

export default LoginButton;