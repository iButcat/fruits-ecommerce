import React, { createContext, useReducer } from 'react';

import Footer from '../components/Footer';
import NavbarI from '../components/Navbar';

function Layout(props) {
    return (
        <div>
            <NavbarI/>
                {props.children}
            <Footer/>
        </div>
    );
}

export default Layout;