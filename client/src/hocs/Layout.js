import React from 'react';
import Footer from '../components/Footer';
import NavbarI from '../components/Navbar';

function Layout(props) {
    return (
        <div>
            <NavbarI isLogged={Boolean(localStorage.getItem('logged'))} />
                {props.children}
            <Footer/>
        </div>
    );
}

export default Layout;