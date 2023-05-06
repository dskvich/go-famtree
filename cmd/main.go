package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-openapi/runtime/middleware"

	"github.com/sushkevichd/go-famtree/pkg/infra/httpserver"

	"github.com/sushkevichd/go-famtree/api/restapi/operations/users"

	"github.com/sushkevichd/go-famtree/pkg/handler"

	"github.com/go-chi/chi"
	"github.com/go-openapi/loads"
	"github.com/sushkevichd/go-famtree/api/restapi"

	"github.com/sushkevichd/go-famtree/api/restapi/operations"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/sushkevichd/go-famtree/config"
	"github.com/sushkevichd/go-famtree/pkg/infra/db"
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
	//treeRepo := db.NewTreeBunRepository(pg)

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Error().Err(err).Msg("Embedded swagger file loading failed")
		os.Exit(1)
	}
	api := operations.NewFamilyTreeAPI(swaggerSpec)

	// Handlers
	userHandler := handler.NewUserHandler(userRepo)
	api.UsersListUsersHandler = users.ListUsersHandlerFunc(userHandler.ListUsers)

	router := chi.NewRouter()
	router.Use(
		httpserver.FileServerMiddleware,
	)

	// Setup swagger UI
	router.Handle("/api/swagger.json", middleware.Spec("/api", swaggerSpec.Raw(), nil))
	router.Handle("/api", middleware.SwaggerUI(middleware.SwaggerUIOpts{
		SpecURL: "/api/swagger.json",
		Path:    "/api",
	}, http.NotFoundHandler()))

	router.Mount("/", api.Serve(nil))

	srv := restapi.NewServer(api)
	defer srv.Shutdown()
	srv.Port = cfg.Port
	srv.ConfigureAPI()
	srv.SetHandler(router)
	if err := srv.Serve(); err != nil {
		log.Error().Err(err).Msg("Server error")
		os.Exit(1)
	}
}
