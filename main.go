package main

import (
	"encoding/json"
	"net/http"

	"github.com/joffrua/go-famtree/internal/httpserver"
)

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// User example
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// GetAllUsers godoc
// @Summary List all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} User
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /users [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{"John", "Doe"},
		{"Mary", "Jane"},
	}
	json.NewEncoder(w).Encode(users)
}

// GetUser godoc
// @Summary Get user by given user_id
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} User
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /users/{user_id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	user := User{"John", "Doe"}
	json.NewEncoder(w).Encode(user)
}

// @title Go Family Tree API
// @version 1.0
// @description Some description
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
func main() {
	s := httpserver.NewBuilder()
	s.AddRoute(http.MethodGet, "/api/users", GetAllUsers)
	s.AddRoute(http.MethodGet, "/api/users/1", GetUser)
	s.AddSwagger("/swagger/")
	s.ServeStatic("/", "./build")

	s.ListenAndServe()
}
