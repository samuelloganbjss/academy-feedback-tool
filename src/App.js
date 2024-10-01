import React, { useEffect, useState } from 'react';
import axios from 'axios';

const App = () => {
  const [students, setStudents] = useState(''); 

  useEffect(() => {
    axios.get('http://localhost:8080/api/students')
      .then(response => {
        console.log(response.data); 
        setStudents(response.data);  
      })
      .catch(error => console.error('Error fetching students:', error));
  }, []);

  return (
    <div>
      <h1>Academy Feedback Tool</h1>
      <p>{students}</p> {}
    </div>
  );
};

export default App;
