import React, { useState } from 'react';
import { auth } from './firebase';

const Signup = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSignup = (e) => {
    e.preventDefault();
    auth.createUserWithEmailAndPassword(email, password)
      .then(userCredential => {
        console.log("User signed up:", userCredential.user);
      })
      .catch(error => console.error(error));
  };

  return (
    <form onSubmit={handleSignup}>
      <input
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        placeholder="Email"
      />
      <input
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        placeholder="Password"
      />
      <button type="submit">Sign Up</button>
    </form>
  );
};

export default Signup;
