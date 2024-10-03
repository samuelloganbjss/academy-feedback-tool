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
      axios.post(`http://localhost:8080/api/students/reports`, {
        student_id: selectedStudent,
        tutorID: 1,
        content: reportContent,
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
    <div className="container mt-5">
      <h1 className="text-center mb-4">Academy Feedback Tool</h1>
  
      <div className="mb-4">
        <h2>Select a Student:</h2>
        <select
          className="form-select mb-3"
          onChange={(e) => {
            const selectedValue = e.target.value;
            setSelectedStudent(selectedValue ? parseInt(selectedValue, 10) : null);
            console.log(selectedValue);
          }}
        >
          <option value="">--Select a student--</option>
          {students.map(student => (
            <option key={student.id} value={student.id}>
              {student.name}
            </option>
          ))}
        </select>
      </div>
  
      <div className="mb-4">
        <h2>Add Report:</h2>
        <textarea
          className="form-control mb-3"
          value={reportContent}
          onChange={(e) => setReportContent(e.target.value)}
          placeholder="Enter report here..."
          rows="5"
        />
        <button onClick={handleReportSubmit} className="btn btn-primary">
          Submit Report
        </button>
      </div>
    </div>
  );
};

export default Home;

