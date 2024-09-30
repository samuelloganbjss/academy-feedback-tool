package main

import (
    "fmt"
    "net/http"
	"io"
)

func rootHandler(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Hello, welcome to the Feedback tool for the academy!")
}

func getTutors(writer http.ResponseWriter, request *http.Request) {
	
    fmt.Printf("got /api/tutors request\n")
	io.WriteString(writer, "a list of all tutors from the db")

}

func getStudents(writer http.ResponseWriter, request *http.Request) {
	
    fmt.Printf("got /api/students  request\n")
	io.WriteString(writer, "a list of all students from the db")

}

func main() {

    http.HandleFunc("/", rootHandler)
	http.HandleFunc("/api/students", getStudents)
	http.HandleFunc("/api/tutors", getTutors)

    fmt.Println("Server listening on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}

