package main

import (
	"feedback-tool/api"
	"feedback-tool/config"
	"feedback-tool/db"
	"feedback-tool/repository/student"
	"feedback-tool/service"
	"fmt"
	"io"
	"net/http"
	"github.com/rs/cors" 
)

func rootHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, welcome to the Feedback tool for the academy!")
}

func getTutors(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("got /api/tutors request\n")
	io.WriteString(writer, "a list of all tutors from the db")
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		// Continue with the next handler
		next.ServeHTTP(writer, request)
	})
}

func getStudents(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("got /api/students  request\n")
	io.WriteString(writer, "a list of all students from the db")
}

func initializeDatabase(config config.DatabaseConfig) (student.StudentRepository, error) {
    switch config.Type {
    case "inmemory":
        return db.NewInMemoryRepository(), nil
    default:
        return nil, fmt.Errorf("unsupported database type: %s", config.Type)
    }
}

func main() {

	// Use the configuration for the in-memory database from Layering-backend
	config := config.InMemory

	// Initialize the database from Layering-backend
	dbRepo, err := initializeDatabase(config)
	if err != nil {
		fmt.Println("Error initializing the database:", err)
		return
	}
	defer dbRepo.Close()

	studentService := service.NewStudentService(dbRepo)
	studentAPI := api.NewStudentAPI(studentService)

	router := http.NewServeMux()

	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/api/students", studentAPI.GetStudents)
	router.HandleFunc("/api/tutors", getTutors)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	handler := c.Handler(router)

	fmt.Println("Server listening on port 8080...")
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
