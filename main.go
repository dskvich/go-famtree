package main

import (
	"context"
	"net/http"

	db "github.com/joffrua/go-famtree/internal/infra/db"

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
	pg := db.NewPg(context.Background())
	userRepo := db.NewUserPgRepository(pg)
	treeRepo := db.NewTreePgRepository(pg)

	userCtrl := controller.NewUserController(userRepo)
	treeCtrl := controller.NewTreeController(treeRepo)

	s := httpserver.NewBuilder()
	s.AddRoute(http.MethodGet, "/api/users", userCtrl.GetAllUsers)
	s.AddRoute(http.MethodGet, "/api/users/{id}", userCtrl.GetUser)

	s.AddRoute(http.MethodPost, "/api/trees", treeCtrl.NewTree)
	s.AddRoute(http.MethodGet, "/api/trees", treeCtrl.GetAllTrees)
	s.AddRoute(http.MethodGet, "/api/trees/{id}", treeCtrl.GetTree)
	s.AddRoute(http.MethodPut, "/api/trees/{id}", treeCtrl.UpdateTree)
	s.AddRoute(http.MethodDelete, "/api/trees/{id}", treeCtrl.DeleteTree)

	s.AddSwagger("/swagger/")
	s.ServeStatic("/", "./build")

	s.ListenAndServe()
}
