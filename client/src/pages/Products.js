import React, { useState, useEffect, useRef } from 'react';

import axios from 'axios';
import { Container, Card, Button, Form, Row, Col } from 'react-bootstrap';

function Products() {
    const [products, setProducts] = useState([]);
    const [productToAdd, setProductToAdd] = useState({
        ID: 0,
        product_name: "",
        price: 0.00,
        quantity: 0,
    });
    const [quantity, setQuantity] = useState(0);
    const quantityRef = useRef(0);

    const addQuantity = (event) => {
        setQuantity(event.target.value);
    };

    const addToCart = (id) => {
        let product = products;
        setProductToAdd({
            ID: product[id].ID,
            product_name: product[id].name,
            price: product[id].price,
            quantity: Number(quantity)
        });
    };

    const getProducts = () => {
        return axios.get('http://localhost:8080/v1/products/getall');
    };

    useEffect(() => {
        getProducts().then((response) => setProducts(response.data.products))
        .catch((error) => console.log(error));
    }, []);

    useEffect(() => {
        var token = localStorage.getItem('token');

        var config = {
            headers: { Authorization: `Bearer ${token}` }
        };
        console.log(productToAdd);
        if (productToAdd.product_name.length !== 0 && quantity !== 0) {
            axios.post('http://localhost:8080/v1/cart/add', productToAdd, config)
            .then(response => console.log(response.data));
        } else {
            return "empty";
        }
    }, [productToAdd]);

    return (
        <Container>
            <div className="products">
                <h1>Products</h1>
                <Row xs={12} sm={12} md={4}>
                {products.length > 0 && products.map((product, id) => (
                    <Col key={id}>
                    <Card style={{ width: '18rem' }}>
                    <Card.Img variant="top" src="holder.js/100px180" />
                    <Card.Body>
                      <Card.Title>{product.name}</Card.Title>
                      <Card.Text>
                            <p>Price: {product.price}</p>
                      </Card.Text>
                      <Form>
                        <Form.Group className="mb-3" controlId="formBasicQuantity">
                            <Form.Control type="quantity" placeholder="Quantity" 
                            ref={quantityRef} 
                            onChange={e => addQuantity(e)} />
                        </Form.Group>
                        <Button 
                        onClick={() => addToCart(id)}
                        variant="primary">Add to cart</Button>
                        </Form>
                    </Card.Body>
                  </Card>
                </Col>
                ))}
                </Row>
            </div>
        </Container>
    );
}

export default Products;