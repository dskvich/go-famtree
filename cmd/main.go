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

func main() {
	log.SetOutput(os.Stdout)

	cfg := new(config.Config)
	if err := envconfig.Process("", cfg); err != nil {
		log.Panicf("config initialization failed: %+v", err)
	}
	log.Infof("config loaded: %+v", cfg)

	pg := db.NewPg(cfg)
	defer pg.Disconnect()

	userRepo := db.NewUserBunRepository(pg)
	treeRepo := db.NewTreeBunRepository(pg)

	userCtrl := controller.NewUserController(userRepo)
	treeCtrl := controller.NewTreeController(treeRepo)

	s := httpserver.NewBuilder(cfg)
	s.AddRoute(http.MethodPost, "/api/users", userCtrl.New)
	s.AddRoute(http.MethodGet, "/api/users", userCtrl.GetAll)
	s.AddRoute(http.MethodGet, "/api/users/{id}", userCtrl.Get)
	s.AddRoute(http.MethodPut, "/api/users/{id}", userCtrl.Update)
	s.AddRoute(http.MethodDelete, "/api/users/{id}", userCtrl.Delete)

	s.AddRoute(http.MethodPost, "/api/trees", treeCtrl.New)
	s.AddRoute(http.MethodGet, "/api/trees", treeCtrl.GetAll)
	s.AddRoute(http.MethodGet, "/api/trees/{id}", treeCtrl.Get)
	s.AddRoute(http.MethodPut, "/api/trees/{id}", treeCtrl.Update)
	s.AddRoute(http.MethodDelete, "/api/trees/{id}", treeCtrl.Delete)

	s.AddStaticDir("/", "./build")

	s.Start()
}
