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
            console.log(studentId)
            const response = await axios.get(`http://localhost:8080/admin/students/reports/${studentId}`, {
                headers: {
                    'Role': 'admin'
                }
            });
            console.log(response.data)
            const reports = response.data;

            createCSV(reports);
        } catch (error) {
            console.error(`Error generating report for student ID ${studentId}:`, error);
        }
    };

    const createCSV = (reports) => {
        const csvRows = [
            ['Content', 'Timestamp']
        ];

        reports.forEach(report => {
            csvRows.push([report.content, report.timestamp]);
        });

        const csvContent = csvRows.map(row => row.join(',')).join('\n');

        const blob = new Blob([csvContent], { type: 'text/csv' });

        const link = document.createElement('a');
        link.href = URL.createObjectURL(blob);
        link.download = 'student_reports.csv';

        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
    };



    return (
        <div className="container mt-5">
            <h1 className="text-center mb-4">Admin Dashboard</h1>

            <h2>Add Student</h2>
            <input
                type="text"
                value={studentName}
                onChange={(e) => setStudentName(e.target.value)}
                placeholder="Enter student name"
                className="form-control mb-2"
            />
            <input
                type="text"
                value={studentDepartment}
                onChange={(e) => setStudentDepartment(e.target.value)}
                placeholder="Enter student department"
                className="form-control mb-2"
            />
            <button onClick={addStudent} className="btn btn-primary">Add Student</button>

            <h2 className="mt-4">Students</h2>
            <table className="table table-striped">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Department</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {students.map((student) => (
                        <tr key={student.id}>
                            <td>{student.name}</td>
                            <td>{student.department}</td>
                            <td>
                                <button onClick={() => toggleViewReports(student.id)} className="btn btn-info btn-sm me-2">
                                    {selectedStudentId === student.id ? 'Hide Reports' : 'View Reports'}
                                </button>
                                <button onClick={() => generateStudentReport(student.id)} className="btn btn-secondary btn-sm me-2">Generate Report</button>
                                <button onClick={() => removeStudent(student.id)} className="btn btn-danger btn-sm">Remove Student</button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>

            {students.map((student) => (
                selectedStudentId === student.id && reports[student.id] && (
                    <div key={student.id} className="mt-3">
                        <h5>Reports for {student.name}</h5>
                        <ul className="list-group">
                            {reports[student.id].length > 0 ? (
                                reports[student.id].map((report) => (
                                    <li key={report.id} className="list-group-item">
                                        {report.content} - {new Date(report.timestamp).toLocaleString()}
                                    </li>
                                ))
                            ) : (
                                <li className="list-group-item">No reports available for this student.</li>
                            )}
                        </ul>
                    </div>
                )
            ))}

            <h2 className="mt-4">Add Tutor</h2>
            <input
                type="text"
                value={tutorName}
                onChange={(e) => setTutorName(e.target.value)}
                placeholder="Enter tutor name"
                className="form-control mb-2"
            />
            <input
                type="text"
                value={tutorDepartment}
                onChange={(e) => setTutorDepartment(e.target.value)}
                placeholder="Enter tutor department"
                className="form-control mb-2"
            />
            <button onClick={addTutor} className="btn btn-primary">Add Tutor</button>

            <h2 className="mt-4">Tutors</h2>
            <table className="table table-striped">
                <thead>
                    <tr>
                        <th>Name</th>
                        <th>Department</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {tutors.map((tutor) => (
                        <tr key={tutor.id}>
                            <td>{tutor.name}</td>
                            <td>{tutor.department}</td>
                            <td>
                                <button onClick={() => removeTutor(tutor.id)} className="btn btn-danger btn-sm">Revoke Access</button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
}

export default AdminDashboard;
