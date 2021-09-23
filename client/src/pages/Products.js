import React, { useState, useEffect } from 'react';

import axios from 'axios';
import { Container, Card, Button, Form } from 'react-bootstrap';

function Products() {
    const [products, setProducts] = useState([]);
    const [productToAdd, setProductToAdd] = useState({
        ID: 0,
        product_name: "",
        price: 0.00,
        quantity: 0,
    });
    const [quantity, setQuantity] = useState(0);

    const addToCart = id => {
        let product = products;
        setProductToAdd({
            ID: product[id].ID,
            product_name: product[id].name,
            price: product[id].price,
            quantity: quantity
        });
        console.log(productToAdd);

        var token = JSON.parse(localStorage.getItem('token'));

        var config = {
            headers: { Authorization: `Bearer ${token}` }
        };
        console.log(productToAdd);
        axios.post('http://localhost:8080/v1/cart/add', productToAdd, config)
        .then(response => console.log(response.data));
        return console.log("done");
    };
    

    const getProducts = () => {
        return axios.get('http://localhost:8080/v1/products/getall');
    };

    useEffect(() => {
        getProducts().then((response) => setProducts(response.data.products))
        .catch((error) => console.log(error));
    }, []);

    return (
        <Container>
            <div className="products">
                <h1>Products</h1>
                {products.length > 0 && products.map((product, id) => (
                    <div key={id}>
                    <Card style={{ width: '18rem' }}>
                    <Card.Img variant="top" src="holder.js/100px180" />
                    <Card.Body>
                      <Card.Title>{product.name}</Card.Title>
                      <Card.Text>
                            <p>Price: {product.price}</p>
                            <p>Quantity: {product.quantity}</p>
                      </Card.Text>
                      <Form>
                        <Form.Group className="mb-3" controlId="formBasicQuantity">
                            <Form.Control type="quantity" placeholder="Quantity" 
                            value={quantity} 
                            onChange={e => setQuantity(e.target.value)} />
                        </Form.Group>
                        <Button 
                        onClick={() => addToCart(id)}
                        variant="primary">Add to cart</Button>
                        </Form>
                    </Card.Body>
                  </Card>
                </div>
                ))}
            </div>
        </Container>
    );
}

export default Products;