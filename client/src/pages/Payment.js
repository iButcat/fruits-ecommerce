import axios from 'axios';
import React, { useState, useEffect } from 'react';

import { Button } from 'react-bootstrap';
import { Link } from 'react-router-dom';

function Payment() {
    const [hasPaid, setHasPaid] = useState(false);
    const shippingAddress = {
        Address: "30 rue de Pierre Fontaine",
        City: "Val de Briey", 
        PostalCode: 54150,
        Country: "France"
    }

    const sendPayment = () => {
        var token = localStorage.getItem('token');

        var config = {
            headers: { Authorization: `Bearer ${token}` }
        };
        let dataPayment = {
            cart_id: localStorage.getItem('cart_id'),
            paid: hasPaid
        };
        return axios.post('http://localhost:8080/v1/payment/create', 
        JSON.stringify(dataPayment), config);
    };

    useEffect(() => {
        sendPayment().then((response) => {
            if (response.status === 200) {
                setHasPaid(true);
            } else {
                return "not paid";
            }
        }).catch((err) => console.log(err));
    }, [hasPaid]);

    return (
        <div className="payment">
            <p>{shippingAddress.Address}</p>
            <p>{shippingAddress.City}</p>
            <p>{shippingAddress.PostalCode}</p>
            <p>{shippingAddress.Country}</p>
            <Button>
            {hasPaid ? 
            <Link to="/home" style={{ textDecoration: 'none' }}>Pay</Link> 
            : 
            ""}
            </Button>
        </div>
    );
}

export default Payment;