import { useEffect, useState } from 'react';

import axios from 'axios';
import { Container } from 'react-bootstrap';
import { Link } from 'react-router-dom';

function Cart() {
    const [cart, setCart] = useState([]);
    const [isLoading, setIsLoading] = useState(true);

    const getCart = () => {
        var token = localStorage.getItem('token');

        var config = {
            headers: { Authorization: `Bearer ${token}` }
        };
        return axios.get('http://localhost:8080/v1/cart/list', config)
    };

    useEffect(() => {
        getCart()
        .then((response) => {
            if (response.status === 200) {
                setCart(response.data.cart);
                setIsLoading(false);
                if (cart.ID !== 0) {
                    localStorage.setItem('cart_id', cart.ID);
                }
                return;
            } else {
                return;
            }
        });
    }, [cart.ID])

    return (
        <Container>
            <div className="cart">
            <h1>{cart.username}</h1>
            <p>Total: {cart.total_price}</p>
            { cart.products ?             
            isLoading !== true && cart.products.map((product, id) => {
                return (
                <div key={id}>
                    <h1>Name: {product.name}</h1>
                    <p>Price: {product.total_price}</p>
                    <p>Quantity: {product.quantity}</p>
                </div>
                );
            }) :
            <h1>Cart is empty</h1>
            }
            <Link to="/payment" className="btn btn-primary">Pay</Link>
            </div>
        </Container>
    );
}

export default Cart;
