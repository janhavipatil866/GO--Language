package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Sample data
var users = []User{
	{ID: 1, Name: "Janhavi", Email: "janhavi@gmail.com"},
	{ID: 2, Name: "Rahul", Email: "rahul@gmail.com"},
}

// Handler function
func getUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Day5 Go HTTP Server")
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/users", getUsers)

	fmt.Println("Server running on port 8080")

	http.ListenAndServe(":8080", nil)
}