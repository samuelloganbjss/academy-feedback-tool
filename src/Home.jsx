import React, { useEffect, useState } from 'react';
import axios from 'axios';

const Home = () => {
  const [students, setStudents] = useState([]);
  const [selectedStudent, setSelectedStudent] = useState(null);
  const [reportContent, setReportContent] = useState('');

  useEffect(() => {
    axios.get('http://localhost:8080/api/students')
      .then(response => {
        setStudents(response.data);
      })
      .catch(error => console.error('Error fetching students:', error));
  }, []);

  const handleReportSubmit = () => {
    if (selectedStudent && reportContent) {
      axios.post(`http://localhost:8080/api/students/reports/${selectedStudent}`, {
        student_id: 1,
        content: reportContent,
        tutorID: 1 
      },
      {
        headers: {
          'Content-Type': 'application/json',
          'Role': 'admin',
        }

      }
    )
      .then(response => {
        console.log(reportContent)
        console.log('Report added:', response.data);
        setReportContent(''); 
      })
      .catch(error => console.error('Error adding report:', error));
    }
  };

  return (
    <div>
      <h1>Academy Feedback Tool</h1>

      <h2>Select a Student:</h2>
      <select onChange={(e) => setSelectedStudent(e.target.value)}>
        <option value="">--Select a student--</option>
        {students.map(student => (
          <option key={student.id} value={student.id}>
            {student.name}
          </option>
        ))}
      </select>

      <h2>Add Report:</h2>
      <textarea
        value={reportContent}
        onChange={(e) => setReportContent(e.target.value)}
        placeholder="Enter report here..."
      />
      <button onClick={handleReportSubmit}>Submit Report</button>
    </div>
  );
};

export default Home;

