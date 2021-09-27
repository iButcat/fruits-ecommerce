import React, { useEffect, useState, refetch } from 'react';

import axios from 'axios';
import { Container, Form, Button } from 'react-bootstrap';
import { Redirect, Link } from 'react-router-dom';

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
                localStorage.setItem('token', response.data.token);
            } else {
                return "wrong login credentials";
            }
            localStorage.setItem('logged', JSON.stringify(isLogged));
            setIsLogged(true);
        })
        .catch((err) => console.log(err));
    }

    useEffect(() => {
        window.addEventListener("storage", () => {
            refetch();
        });

        console.log('UPDATE LOGGED: ', isLogged);
        localStorage.setItem('logged', JSON.stringify(isLogged));
    }, [isLogged]);

    return (
        <Container style={{padding: 20}}>
            <div className="login">
                <Form onSubmit={handleSubmit}>
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
                    {isLogged ? <Redirect to="/home" style={{ textDecoration: 'none' }}>Submit</Redirect> : <Link/>}
                </Button>
                </Form>
            </div>
        </Container>
    );
}

export default Login;