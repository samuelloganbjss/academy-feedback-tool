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

func initializeDatabase(config config.DatabaseConfig) (student.StudentRepository, error) {
    switch config.Type {
    // case "postgres":
    //     connectionString := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", config.User, config.DBName, config.Password, config.Host, config.SSLMode)
    //     return postgres.NewPostgresRepository(connectionString)
    case "inmemory":
        return db.NewInMemoryRepository(), nil
    default:
        return nil, fmt.Errorf("unsupported database type: %s", config.Type)
    }
}
func main() {

	config := config.InMemory

	dbRepo, erro := initializeDatabase(config)
    if erro != nil {
        fmt.Println("Error initializing the database:", erro)
        return
    }
    defer dbRepo.Close()

	studentService := service.NewStudentService(dbRepo)
    studentAPI := api.NewStudentAPI(studentService)
	
	router := http.NewServeMux()

	router.HandleFunc("GET /", rootHandler)
	router.HandleFunc("GET /api/students", studentAPI.GetStudents)
	router.HandleFunc("GET /api/tutors", getTutors)

	fmt.Println("Server listening on port 8080...")

	err := http.ListenAndServe(":8080", CorsMiddleware(router))

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
