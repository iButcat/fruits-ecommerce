import React from 'react';

import { Redirect } from 'react-router-dom';
 
function PaymentButton() {
    return <Redirect to="/payment" style={{ textDecoration: 'none' }}>Pay</Redirect>;
}

export default PaymentButton;