import React, { useState } from 'react';

import axios from 'axios';

function Register() {
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const handleSubmit = (events) => {
        events.preventDefault();
        axios.post(
            'http://localhost:8080/v1/user/register', JSON.stringify({
                email: email,
                username: username, 
                password: password,
            }))
        .then(response => console.log(response.data))
        .catch((err) => console.log(err));
    }

    return (
        <div className="register">
            <h1>Register</h1>
            <form onSubmit={handleSubmit}>
                <label>
                    <input 
                    placeholder="Email" 
                    type="text" value={email} 
                    onChange={e => setEmail(e.target.value)} />
                </label>
                <label>
                    <input 
                    placeholder="Username" 
                    type="text" value={username} 
                    onChange={e => setUsername(e.target.value)} />
                </label>
                <label>
                    <input 
                    placeHolder="Password" 
                    type="password" value={password} 
                    onChange={e => setPassword(e.target.value)} />
                </label>
                <input type="submit" value="Submit" />
            </form>
        </div>
    );
}

export default Register;