import React from 'react';
import Footer from '../components/Footer';
import Navbar from '../components/Navbar';

function Layout(props) {
    return (
        <div>
            <Navbar/>
                {props.children}
            <Footer/>
        </div>
    );
}

export default Layout;