import React, { useState } from 'react';

import axios from 'axios';
import { Container, Form, Button } from 'react-bootstrap';
import { Link } from 'react-router-dom';

function Login() {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [isLogged, setIsLogged] = useState(false);

    const handleSubmit = (events) => {
        events.preventDefault();
        axios.post(
            'http://localhost:8080/v1/user/login', JSON.stringify({
                username: username, 
                password: password,
            }))
        .then(response => {
            if (response.status === 200) {
                setIsLogged(true);
                localStorage.setItem('logged', JSON.stringify(isLogged));
                localStorage.setItem('token', JSON.stringify(response.data.token));
            } else {
                return "wrong login credentials";
            }
        })
        .catch((err) => console.log(err));
    }
    return (
        <Container>
            <div className="login">
                {isLogged 
                ? <h1>Already logged in</h1>
                : <Form onSubmit={handleSubmit}>
                <Form.Group className="mb-3" controlId="formBasicUsername">
                    <Form.Label>Username</Form.Label>
                    <Form.Control type="username" placeholder="Enter username" 
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
                    {isLogged ? <Link to="home" /> : <Link/>}
                </Button>
                </Form>
                }
            </div>
        </Container>
    );
}

export default Login;