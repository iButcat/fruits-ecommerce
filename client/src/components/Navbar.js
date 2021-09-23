import React, { useState } from 'react';
import  { Navbar, Nav, Container } from 'react-bootstrap';

import LogoutButton from '../components/LogoutButton';

function NavbarI(props) {
    const [logged, setLogged] = useState(false); 
    
    console.log("PROPS LOGGED", props.isLogged);

    window.addEventListener("storage", (e) => {
        setLogged(Boolean(localStorage.getItem('logged')));
    });

    return (
        <Navbar bg="light" expand="lg">
            <Container> 
                <Navbar.Brand href="home">Cinemo Fruits</Navbar.Brand>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse id="basic-navbar-nav">
                <Nav className="me-auto">
                    {props.isLogged === true ? 
                    <Nav.Link><LogoutButton/></Nav.Link> :
                    <div> 
                    <Nav.Link href="login">Login</Nav.Link>
                    <Nav.Link href="register">Register</Nav.Link>
                    </div>
                    }
                </Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
}

export default NavbarI;