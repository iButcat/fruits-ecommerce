import React, { useState } from 'react';

import axios from 'axios';

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
        <div className="login">
            {isLogged 
            ? <h1>Already logged in</h1>
            : <form onSubmit={handleSubmit}>
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
            }
        </div>
    );
}

export default Login;