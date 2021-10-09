package main

import (
	"net/http"

	"github.com/joffrua/go-famtree/internal/controller"

	"github.com/joffrua/go-famtree/internal/infra/httpserver"
)

// @title Go Family Tree API
// @version 1.0
// @description Some description
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api
func main() {
	userCtrl := controller.NewUserController()
	treeCtrl := controller.NewTreeController()

	s := httpserver.NewBuilder()
	s.AddRoute(http.MethodGet, "/api/users", userCtrl.GetAllUsers)
	s.AddRoute(http.MethodGet, "/api/users/{userID}", userCtrl.GetUser)

	s.AddRoute(http.MethodPost, "/api/trees", treeCtrl.NewTree)
	s.AddRoute(http.MethodGet, "/api/trees", treeCtrl.GetAllTrees)
	s.AddRoute(http.MethodGet, "/api/trees/{treeID}", treeCtrl.GetTree)
	s.AddRoute(http.MethodPatch, "/api/trees/{treeID}", treeCtrl.UpdateTree)
	s.AddRoute(http.MethodDelete, "/api/trees/{treeID}", treeCtrl.DeleteTree)

	s.AddSwagger("/swagger/")
	s.ServeStatic("/", "./build")

	s.ListenAndServe()
}
