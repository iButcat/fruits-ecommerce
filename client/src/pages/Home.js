import React from 'react';
import { Container, Row, Col } from 'react-bootstrap';

function Home() {
    return (
        <Container className="home" style={{padding: 20}}>
           <Row lg={5}>
               <Col lg={12} md={8}>
                   <h3 className="text-center">Cinemo Fruits Ecommerce</h3>
                   <p>Buy the best fruits with some 
                       great discout depending on the quantity</p>
               </Col>
           </Row>
        </Container>
    );
}

export default Home;