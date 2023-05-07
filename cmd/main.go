package main

import (
	"net/http"
	"os"
	"time"

	"github.com/sushkevichd/go-famtree/pkg/infrastructure/database"

	"github.com/sushkevichd/go-famtree/api/restapi/operations/users"

	"github.com/sushkevichd/go-famtree/pkg/handler"

	"github.com/go-chi/chi"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sushkevichd/go-famtree/api/restapi"

	"github.com/sushkevichd/go-famtree/api/restapi/operations"

	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/sushkevichd/go-famtree/config"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})

	cfg := new(config.Config)
	if err := envconfig.Process("", cfg); err != nil {
		log.Error().Err(err).Msg("initializing config")
		os.Exit(1)
	}
	log.Info().Msgf("Config loaded: %+v", cfg)

	postgres, err := database.NewPostgres(cfg.PG.URL, cfg.PG.Host)
	if err != nil {
		log.Error().Err(err).Msg("initializing PostgreSQL connection: %v")
		os.Exit(1)
	}
	defer database.ClosePostgres(postgres)

	bunDB := database.NewBunDB(postgres, cfg.PG.ShowSQL)

	userRepo := database.NewUserBunRepository(bunDB)
	//treeRepo := db.NewTreeBunRepository(postgres)

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Error().Err(err).Msg("Embedded swagger file loading failed")
		os.Exit(1)
	}
	api := operations.NewFamilyTreeAPI(swaggerSpec)

	// Handlers
	userHandler := handler.NewUserHandler(userRepo)
	api.UsersGetUsersHandler = users.GetUsersHandlerFunc(userHandler.GetUsers)
	api.UsersGetUserByIDHandler = users.GetUserByIDHandlerFunc(userHandler.GetUserByID)
	api.UsersCreateUserHandler = users.CreateUserHandlerFunc(userHandler.CreateUser)
	api.UsersUpdateUserByIDHandler = users.UpdateUserByIDHandlerFunc(userHandler.UpdateUserByID)
	api.UsersDeleteUserByIDHandler = users.DeleteUserByIDHandlerFunc(userHandler.DeleteUserByID)

	router := chi.NewRouter()
	router.Use(
		handler.FileServerMiddleware,
		//handler.LoggerMiddleware,
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
