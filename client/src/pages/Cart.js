import React, { useEffect, useState } from 'react';

import axios from 'axios';
import { Container } from 'react-bootstrap';

function Cart() {
    const [cart, setCart] = useState([]);
    const [isLoading, setIsLoading] = useState(false);

    const getCart = () => {
        var token = JSON.parse(localStorage.getItem('token'));

        var config = {
            headers: { Authorization: `Bearer ${token}` }
        };
        return axios.get('http://localhost:8080/v1/cart/list', config);
    };

    useEffect(() => {
        getCart().then((response) => setCart(response.data.cart))
        .catch((error) => console.log(error));
        if (cart.length !== 0) {
            setIsLoading(true);
        }
    }, []);


    return (
        <Container>
            <div className="cart">
            <h1>{cart.username}</h1>
            {isLoading && cart.products.length > 0 && cart.products.foreach((product, id) => {
                {console.log(product)}
                <div key={id}>
                    <h1>{product.ID}</h1>
                    <h1>{product.name}</h1>
                    <p>{product.price}</p>
                </div>
            })}
            </div>
        </Container>
    );
}

export default Cart;