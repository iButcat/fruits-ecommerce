import React, { useState, useEffect } from 'react';

import axios from 'axios';

function Products() {
    const [products, setProducts] = useState([]);

    const getProducts = () => {
        return axios.get('http://localhost:8080/v1/products/getall');
    };

    useEffect(() => {
        getProducts().then((response) => setProducts(response.data.products))
        .catch((error) => console.log(error));
    }, []);

    return (
        <div className="products">
            {products.length > 0 && products.map(product => (
                <div key={product.ID}>
                    <h1>{product.name}</h1>
                    <p>
                        {product.price},
                        {product.quantity}
                    </p>
                </div>
            ))}
        </div>
    );
}

export default Products;