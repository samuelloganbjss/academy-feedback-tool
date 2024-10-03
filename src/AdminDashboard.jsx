import React, { useState, useEffect } from 'react';
import axios from 'axios';

const AdminDashboard = () => {
    const [students, setStudents] = useState([]);
    const [tutors, setTutors] = useState([]);
    const [studentName, setStudentName] = useState('');
    const [studentDepartment, setStudentDepartment] = useState('');
    const [tutorName, setTutorName] = useState('');
    const [tutorDepartment, setTutorDepartment] = useState('');
    const [reports, setReports] = useState({});
    const [selectedStudentId, setSelectedStudentId] = useState(null);
    const [fullReport, setFullReport] = useState(null);  


    useEffect(() => {
        fetchStudents();
        fetchTutors();
    }, []);

    const fetchStudents = async () => {
        axios.get('http://localhost:8080/api/students')
            .then(response => {
                setStudents(response.data);
            })
            .catch(error => console.error('Error fetching students:', error));
    };

    const fetchTutors = async () => {
        axios.get('http://localhost:8080/api/tutors')
            .then(response => {
                setTutors(response.data);
            })
            .catch(error => console.error('Error fetching tutors:', error));
    };

    const fetchReports = async (studentId) => {
        console.log(`Fetching reports for student ID: ${studentId}`);
        axios.get(`http://localhost:8080/admin/students/reports/${studentId}`, {
            headers: {
                'Role': 'admin'
            }
        })
        .then(response => {
            console.log('Reports fetched:', response.data);
            setReports((prevReports) => ({
                ...prevReports,
                [studentId]: response.data,
            }));
        })
        .catch(error => console.error('Error fetching reports:', error));
    };

    const addStudent = async () => {
        if (studentName && studentDepartment) {
            axios.post('http://localhost:8080/api/students', {
                name: studentName,
                department: studentDepartment
            }, {
                headers: {
                    'Content-Type': 'application/json',
                    'Role': 'admin',
                }
            })
            .then(response => {
                console.log('student added:', response.data);
                setStudentName('');
                setStudentDepartment('');
                fetchStudents();
            })
            .catch(error => console.error('Error adding student:', error));
        }
    };

    const addTutor = async () => {
        if (tutorName && tutorDepartment) {
            axios.post('http://localhost:8080/api/tutors', {
                name: tutorName,
                department: tutorDepartment
            }, {
                headers: {
                    'Content-Type': 'application/json',
                    'Role': 'admin',
                }
            })
            .then(response => {
                console.log('Tutor added:', response.data);
                setTutorName('');
                setTutorDepartment('');
                fetchTutors();
            })
            .catch(error => console.error('Error adding tutor:', error));
        }
    };

    const removeStudent = async (id) => {
        axios.delete(`http://localhost:8080/api/students/remove/${id}`, {
            headers: {
                'Role': 'admin',
            }
        })
        .then(response => {
            console.log('Student removed:', response.data);
            fetchStudents();
        })
        .catch(error => console.error('Error removing student:', error));
    };

    const removeTutor = async (id) => {
        axios.delete(`http://localhost:8080/api/tutors/remove/${id}`, {
            headers: {
                'Role': 'admin',
            }
        })
        .then(response => {
            console.log('Tutor removed:', response.data);
            fetchTutors();
        })
        .catch(error => console.error('Error removing tutor:', error));
    };

    const toggleViewReports = (studentId) => {
        if (selectedStudentId === studentId) {
            setSelectedStudentId(null); 
        } else {
            setSelectedStudentId(studentId); 
            fetchReports(studentId); 
        }
    };
       
    const generateStudentReport = async (studentId) => {
        try {
            const response = await axios.get(`http://localhost:8080/admin/students/reports?student_id=${studentId}`, {
                headers: {
                    'Role': 'admin'
                }
            });
            setSelectedStudentId(response.data);  
        } catch (error) {
            console.error(`Error generating report for student ID ${studentId}:`, error);
        }
    };

    return (
        <div>
            <h1>Admin Dashboard</h1>

            <h2>Add Student</h2>
            <input
                type="text"
                value={studentName}
                onChange={(e) => setStudentName(e.target.value)}
                placeholder="Enter student name"
            />
            <input
                type="text"
                value={studentDepartment}
                onChange={(e) => setStudentDepartment(e.target.value)}
                placeholder="Enter student department"
            />
            <button onClick={addStudent}>Add Student</button>

            <h2>Students</h2>
            <ul>
            {students.map((student) => (
                <li key={student.id}>
                    {student.name} ({student.department})
                    <button onClick={() => removeStudent(student.id)}>Remove</button>
                    <button onClick={() => toggleViewReports(student.id)}>
                        {selectedStudentId === student.id ? 'Hide Reports' : 'View Reports'}
                    </button>

                    {}
                    <button onClick={() => generateStudentReport(student)}>Generate Report</button>

                    {}
                    {selectedStudentId === student.id && reports[student.id] && (
                        <ul>
                            {reports[student.id].length > 0 ? (
                                reports[student.id].map((report) => (
                                    <li key={report.id}>
                                        {report.content} - {new Date(report.timestamp).toLocaleString()}
                                    </li>
                                ))
                            ) : (
                                <li>No reports available for this student.</li>
                            )}
                        </ul>
                    )}
                </li>
            ))}
        </ul>


            <h2>Add Tutor</h2>
            <input
                type="text"
                value={tutorName}
                onChange={(e) => setTutorName(e.target.value)}
                placeholder="Enter tutor name"
            />
            <input
                type="text"
                value={tutorDepartment}
                onChange={(e) => setTutorDepartment(e.target.value)}
                placeholder="Enter tutor department"
            />
            <button onClick={addTutor}>Add Tutor</button>

            <h2>Tutors</h2>
            <ul>
                {tutors.map((tutor) => (
                    <li key={tutor.id}>
                        {tutor.name} ({tutor.department})
                        <button onClick={() => removeTutor(tutor.id)}>Remove</button>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default AdminDashboard;
