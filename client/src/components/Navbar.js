import React from 'react';

import  { Navbar, Nav, Container } from 'react-bootstrap';

function NavbarI() {
    return (
        <Navbar bg="light" expand="lg">
            <Container> 
                <Navbar.Brand href="home">Cinemo Fruits</Navbar.Brand>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse id="basic-navbar-nav">
                <Nav className="me-auto">
                    <Nav.Link href="login">Login</Nav.Link>
                    <Nav.Link href="register">Register</Nav.Link>
                </Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
}

export default NavbarI;