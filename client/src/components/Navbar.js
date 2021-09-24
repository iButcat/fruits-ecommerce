import React from 'react';
import  { Navbar, Nav, Container } from 'react-bootstrap';

import LogoutButton from '../components/LogoutButton';
import getValueLocalStorage from '../utils/Utils';

function NavbarI() {
    
    return (
        <Navbar bg="light" expand="lg">
            <Container> 
                <Navbar.Brand href="home">Cinemo Fruits</Navbar.Brand>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse id="basic-navbar-nav">
                <Nav className="me-auto">
                    {getValueLocalStorage('logged') === "true" ? 
                    <Nav.Link><LogoutButton/></Nav.Link> :
                    <>
                    <Nav.Link href="login">Login</Nav.Link>
                    <Nav.Link href="register">Register</Nav.Link>
                    </>
                    }
                    <Nav.Link href="cart">Cart</Nav.Link>
                    <Nav.Link href="products">Products</Nav.Link>
                </Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
}

export default NavbarI;