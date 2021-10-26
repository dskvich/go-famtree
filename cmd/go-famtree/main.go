package main

import (
	"net/http"
	"os"

	"github.com/kelseyhightower/envconfig"

	"github.com/joffrua/go-famtree/config"

	log "github.com/sirupsen/logrus"

	"github.com/joffrua/go-famtree/internal/infra/db"

	"github.com/joffrua/go-famtree/internal/controller"

	"github.com/joffrua/go-famtree/internal/infra/httpserver"
)

// Go Family Tree API
//
// Some description
//
//     Schemes: http
//     BasePath: /api
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
func main() {
	log.SetOutput(os.Stdout)

	cfg := new(config.Config)
	if err := envconfig.Process("", cfg); err != nil {
		log.Panicf("config initialization failed: %+v", err)
	}
	log.Infof("config loaded: %+v", cfg)

	pg := db.NewPg(cfg)
	userRepo := db.NewUserPgRepository(pg)
	treeRepo := db.NewTreePgRepository(pg)

	userCtrl := controller.NewUserController(userRepo)
	treeCtrl := controller.NewTreeController(treeRepo)

	s := httpserver.NewBuilder(cfg)
	s.AddRoute(http.MethodGet, "/api/users", userCtrl.GetAllUsers)
	s.AddRoute(http.MethodGet, "/api/users/{id}", userCtrl.GetUser)

	s.AddRoute(http.MethodPost, "/api/trees", treeCtrl.NewTree)
	s.AddRoute(http.MethodGet, "/api/trees", treeCtrl.GetAllTrees)
	s.AddRoute(http.MethodGet, "/api/trees/{id}", treeCtrl.GetTree)
	s.AddRoute(http.MethodPut, "/api/trees/{id}", treeCtrl.UpdateTree)
	s.AddRoute(http.MethodDelete, "/api/trees/{id}", treeCtrl.DeleteTree)

	s.AddStaticDir("/", "./build")

	s.Start()
}
