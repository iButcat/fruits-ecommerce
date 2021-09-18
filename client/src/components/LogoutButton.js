import React from 'react';
import { Link } from 'react-router-dom';

function LogoutButton() {
    
    const handleOnClick = (event) => {
        event.preventDefault();
        localStorage.removeItem('token');
        localStorage.setItem('logged', JSON.stringify(false));
    };

    return (
        <button onClick={handleOnClick}>
            <Link to='/home'>Logout</Link>
        </button>
    );
}

export default LogoutButton;