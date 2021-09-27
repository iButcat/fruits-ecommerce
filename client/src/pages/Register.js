import React, { useState } from 'react';

import axios from 'axios';
import { Container, Form, Button } from 'react-bootstrap';
import { Redirect } from 'react-router-dom';
 
function Register() {
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [registred, setRegistred] = useState(null);

    const clearWhenSubmit = () => {
        setUsername("");
        setPassword("");
        setEmail("");
    };

    const handleSubmit = (events) => {
        events.preventDefault();
        axios.post(
            'http://localhost:8080/v1/user/register', JSON.stringify({
                email: email,
                username: username, 
                password: password,
            }))
        .then(response => {
            if (response.status === 201) {
                setRegistred(registred => true);
                clearWhenSubmit();
                return registred
            } else {
                setRegistred(registred => false);
                return registred
            }
        })
        .catch((err) => console.log(err));
    }

    return (
        <Container>
            <Form onSubmit={handleSubmit} id="register-form">
                <Form.Group className="mb-3" controlId="formBasicEmail">
                    <Form.Label>Email address</Form.Label>
                    <Form.Control 
                    type="email" 
                    placeholder="Enter email" 
                    value={email} 
                    onChange={e => setEmail(e.target.value)} />
                    <Form.Text className="text-muted">
                    We'll never share your email with anyone else.
                    </Form.Text>
                </Form.Group>
                <Form.Group className="mb-3" controlId="formBasicUsername">
                    <Form.Label>Username</Form.Label>
                    <Form.Control 
                    type="username" 
                    placeholder="Enter username" 
                    value={username} 
                    onChange={e => setUsername(e.target.value)} />
                </Form.Group>
                <Form.Group className="mb-3" controlId="formBasicPassword">
                    <Form.Label>Password</Form.Label>
                    <Form.Control 
                    type="password" 
                    placeholder="Password" 
                    value={password} 
                    onChange={e => setPassword(e.target.value)} />
                </Form.Group>
                <Button variant="primary" type="submit">
                Submit
                </Button>
                    {registred ? <Redirect to="login" /> : ""}
            </Form>
        </Container>
    );
}

export default Register;