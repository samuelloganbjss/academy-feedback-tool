package main

import (
	"fmt"
	"net/http"
	"github.com/rs/cors"
	middleware "github.com/samuelloganbjss/academy-feedback-tool/admin"
	"github.com/samuelloganbjss/academy-feedback-tool/api"
	"github.com/samuelloganbjss/academy-feedback-tool/config"
	"github.com/samuelloganbjss/academy-feedback-tool/db"
	"github.com/samuelloganbjss/academy-feedback-tool/repository/student"
	"github.com/samuelloganbjss/academy-feedback-tool/repository/tutor"
	"github.com/samuelloganbjss/academy-feedback-tool/service"
)

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, welcome to the Feedback tool for the academy!")
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(writer, request)
	})
}

func initializeDatabase(config config.DatabaseConfig) (student.StudentRepository, tutor.TutorRepository, error) {
	switch config.Type {
	case "inmemory":
		return db.NewInMemoryRepository(), db.NewInMemoryRepository(), nil
	default:
		return nil, nil, fmt.Errorf("unsupported database type: %s", config.Type)
	}
}

func getTutorRoleFromRequest(r *http.Request) (string, error) {
	role := r.Header.Get("Role")

	if role == "" {
		return "", fmt.Errorf("role not found")
	}
	return role, nil
}

func main() {

	config := config.InMemory

	dbRepoStudent, dbRepoTutor, err := initializeDatabase(config)
	if err != nil {
		fmt.Println("Error initializing the database:", err)
		return
	}
	defer dbRepoStudent.Close()

	studentService := service.NewStudentService(dbRepoStudent)
	studentAPI := api.NewStudentAPI(studentService)

	router := http.NewServeMux()

	router.HandleFunc("/", rootHandler)

	router.HandleFunc("GET /api/students", studentAPI.GetStudents)
	router.Handle("POST /api/students", middleware.AdminMiddleware(getTutorRoleFromRequest)(http.HandlerFunc(studentAPI.AddStudent)))
	router.Handle("DELETE /api/students/remove/{id}", middleware.AdminMiddleware(getTutorRoleFromRequest)(http.HandlerFunc(studentAPI.DeleteSingleStudent)))

	tutorService := service.NewTutorService(dbRepoTutor)
	tutorAPI := api.NewTutorAPI(tutorService)

	router.HandleFunc("GET /api/tutors", tutorAPI.GetTutors)
	router.Handle("POST /api/tutors", middleware.AdminMiddleware(getTutorRoleFromRequest)(http.HandlerFunc(tutorAPI.AddTutor)))
	router.Handle("DELETE /api/tutors/remove/{id}", middleware.AdminMiddleware(getTutorRoleFromRequest)(http.HandlerFunc(tutorAPI.DeleteSingleTutor)))
	router.Handle("POST /api/students/reports", middleware.AdminMiddleware(getTutorRoleFromRequest)(http.HandlerFunc(tutorAPI.AddReport)))
	router.Handle("POST /api/students/reports/edit/{id}", middleware.AdminMiddleware(getTutorRoleFromRequest)(http.HandlerFunc(tutorAPI.EditReport)))
	router.Handle("GET /admin/students/reports/{id}", middleware.AdminMiddleware(getTutorRoleFromRequest)(http.HandlerFunc(tutorAPI.GetStudentReports)))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Role"},
	})

	handler := c.Handler(router)

	fmt.Println("Server listening on port 8080...")
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
