import React, { useEffect, useState } from 'react';

import axios from 'axios';
import { Container } from 'react-bootstrap';

function Cart() {
    const [cart, setCart] = useState([]);

    const getCart = () => {
        var token = JSON.parse(localStorage.getItem('token'));

        console.log(token);
        var config = {
            headers: { Authorization: `Bearer ${token}` }
        };
        return axios.get('http://localhost:8080/v1/cart/list', config);
    };

    useEffect(() => {
        getCart().then((response) => setCart(response.data.cart))
        .catch((error) => console.log(error));
    }, []);

    return (
        <Container>
            <div className="cart">
            {cart.length > 0 && cart.map(cart => (
                <div key={cart.ID}>
                    <h1>{cart.products}</h1>
                    <p>{cart.quantity}</p>
                </div>
            ))}
            </div>
        </Container>
    );
}

export default Cart;