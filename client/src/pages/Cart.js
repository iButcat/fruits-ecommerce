import React, { useEffect, useState } from 'react';

import axios from 'axios';
import { Container } from 'react-bootstrap';

function Cart() {
    const [cart, setCart] = useState([]);
    const [products, setProducts] = useState([]);
    const [isLoading, setIsLoading] = useState(false);

    const getCart = () => {
        var token = localStorage.getItem('token');

        var config = {
            headers: { Authorization: `Bearer ${token}` }
        };
        axios.get('http://localhost:8080/v1/cart/list', config)
        .then((response) => setCart(response.data.cart))
        .catch((error) => console.log(error));
        if (cart.length !== 0) {
            setIsLoading(true);
        } else {
            return;
        }
        setProducts(products => cart.products);
        getCart();
    };

    useEffect(() => {
        console.log('UPDATE PRODUCTS: ', products)
    }, [products]);


    return (
        <Container>
            <div className="cart">
            <h1>{cart.username}</h1>
            {products.length !== null && products.length !== 0 ?             
            isLoading && products.map((product, id) => {
                return (
                <div key={id}>
                    <h1>Name: {product.name}</h1>
                    <p>Quantity: {product.quantity}</p>
                    <p>Price: {product.price}</p>
                </div>
                );
            }):
            <h1>Cart is empty</h1>
            }
            </div>
        </Container>
    );
}

export default Cart;
