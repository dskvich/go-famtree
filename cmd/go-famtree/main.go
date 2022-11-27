package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joffrua/go-famtree/internal/repository"

	"github.com/kelseyhightower/envconfig"

	"github.com/joffrua/go-famtree/config"

	"github.com/sirupsen/logrus"

	"github.com/joffrua/go-famtree/internal/controller"

	"github.com/joffrua/go-famtree/internal/infrastructure/rest"
)

var appName = "go-famtree"

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
	logger := NewLogger()

	cfg := new(config.Config)
	if err := envconfig.Process("", cfg); err != nil {
		logger.WithError(err).Error("env config initialization failed")
		_ = os.Stderr.Sync()
		os.Exit(1)
	}
	logger.WithField("cfg", fmt.Sprintf("%+v", cfg)).Info("env config loaded")

	//TODO: move to a different place
	logLevel, err := logrus.ParseLevel(cfg.Log.Level)
	if err != nil {
		logger.WithError(err).Error("config initialization failed")
	} else {
		logger.Logger.SetLevel(logLevel)
	}

	pg := repository.NewPg(cfg)
	userRepo := repository.NewUserPgRepository(pg)
	treeRepo := repository.NewTreePgRepository(pg)

	userCtrl := controller.NewUserController(userRepo)
	treeCtrl := controller.NewTreeController(treeRepo)

	s := rest.NewService(rest.Config{
		ListenAddr: fmt.Sprintf(":%s", cfg.Port),
		Logger:     logger,
	})
	s.Get("/api/users", userCtrl.GetAllUsers)
	s.Get("/api/users/{id}", userCtrl.GetUser)

	s.Post("/api/trees", treeCtrl.NewTree)
	//s.Get("/api/trees", treeCtrl.GetAllTrees)
	//s.Get("/api/trees/{id}", treeCtrl.GetTree)
	//s.Put("/api/trees/{id}", treeCtrl.UpdateTree)
	//s.Delete("/api/trees/{id}", treeCtrl.DeleteTree)

	//s.AddStaticDir("/", "./build")

	err = s.Run(context.Background())
	if err != nil {
		logger.WithError(err).Error("shutting down due to error")
		_ = os.Stderr.Sync()
		os.Exit(1)
	}
}

func NewLogger() *logrus.Entry {
	rootLogger := logrus.New()
	rootLogger.SetOutput(os.Stdout)
	rootLogger.SetLevel(logrus.InfoLevel)
	rootLogger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05.999",
		FullTimestamp:   true,
	})

	return rootLogger.WithFields(logrus.Fields{
		"app": appName,
	})
}
