package main

import (
	"encoding/json"
	"net/http"

	"github.com/joffrua/go-famtree/internal/httpserver"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{"John", "Doe"},
		{"Mary", "Jane"},
	}
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := User{"John", "Doe"}
	json.NewEncoder(w).Encode(user)
}

func main() {
	s := httpserver.NewBuilder()
	s.AddRoute(http.MethodGet, "/api/users", GetAllUsers)
	s.AddRoute(http.MethodGet, "/api/users/1", GetUser)

	s.ServeStatic("/", "./build")

	s.ListenAndServe()
}
