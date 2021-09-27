import { useEffect, useState } from 'react';

import axios from 'axios';
import { Container } from 'react-bootstrap';

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
                return;
            } else {
                return;
            }
        });
    }, [])

    return (
        <Container>
            <div className="cart">
            <h1>{cart.ID}</h1>
            <h1>{cart.username}</h1>
            { cart.products ?             
            isLoading !== true && cart.products.map((product, id) => {
                return (
                <div key={id}>
                    <h1>Name: {product.name}</h1>
                    <p>Price: {product.price}</p>
                </div>
                );
            }) :
            <h1>Cart is empty</h1>
            }
            </div>
        </Container>
    );
}

export default Cart;
