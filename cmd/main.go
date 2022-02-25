package main

import (
	"net/http"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/joffrua/go-famtree/config"
	"github.com/joffrua/go-famtree/internal/controller"
	"github.com/joffrua/go-famtree/internal/infra/db"
	"github.com/joffrua/go-famtree/internal/infra/httpserver"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})

	cfg := new(config.Config)
	if err := envconfig.Process("", cfg); err != nil {
		log.Error().Err(err).Msg("Config initialization failed")
		os.Exit(1)
	}
	log.Info().Msgf("Config loaded: %+v", cfg)

	pg, err := db.NewPg(&cfg.PG)
	if err != nil {
		log.Error().Err(err).Msg("PG connection error")
		os.Exit(1)
	}
	defer pg.Disconnect()

	userRepo := db.NewUserBunRepository(pg)
	treeRepo := db.NewTreeBunRepository(pg)

	userCtrl := controller.NewUserController(userRepo)
	treeCtrl := controller.NewTreeController(treeRepo)

	s := httpserver.NewBuilder(&cfg.HTTP)
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
