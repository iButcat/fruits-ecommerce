import React, { useState, useEffect } from 'react';

import axios from 'axios';
import { Container, Card, Button } from 'react-bootstrap';

function Products() {
    const [products, setProducts] = useState([]);;
    const [productToAdd, setProductToAdd] = useState({
        name: "",
        quantityCard: 0,
    });

    const incrementQuantity = () => {}

    const addToCart = (e, i) => {
        const product = products;
        console.log("product name: ", product[i].name);
        setProductToAdd({
            id: i,
            name: product[i].name,
            quantityCard: 0
        });
        return productToAdd;
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
                      <Button 
                      variant="primary"
                      onClick={(e) => addToCart(e,id)}>
                          {product.quantityCard}
                        </Button>
                      <Button variant="primary" 
                      onClick={(e) => incrementQuantity(e, id)}>Add to cart</Button>
                    </Card.Body>
                  </Card>
                </div>
                ))}
            </div>
        </Container>
    );
}

export default Products;