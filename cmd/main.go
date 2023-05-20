package main

import (
	"net/http"
	"os"
	"time"

	"github.com/sushkevichd/go-famtree/api/restapi/operations/people"

	"github.com/sushkevichd/go-famtree/api/restapi/operations/trees"

	"github.com/sushkevichd/go-famtree/pkg/infrastructure/database"

	"github.com/sushkevichd/go-famtree/api/restapi/operations/users"

	"github.com/sushkevichd/go-famtree/pkg/handler"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/sushkevichd/go-famtree/api/restapi"

	"github.com/sushkevichd/go-famtree/api/restapi/operations"

	"github.com/go-chi/cors"
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
	treeRepo := database.NewTreeBunRepository(bunDB)
	peopleRepo := database.NewPeopleBunRepository(bunDB)

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

	treeHandler := handler.NewTreeHandler(treeRepo)
	api.TreesGetAllTreesForUserHandler = trees.GetAllTreesForUserHandlerFunc(treeHandler.GetAllTreesForUser)
	api.TreesCreateTreeForUserHandler = trees.CreateTreeForUserHandlerFunc(treeHandler.CreateTreeForUser)

	peopleHandler := handler.NewPeopleHandler(peopleRepo)
	api.PeopleGetPeopleByTreeHandler = people.GetPeopleByTreeHandlerFunc(peopleHandler.GetPeopleByTree)

	router := chi.NewRouter()

	// Add CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // replace with your allowed origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})
	router.Use(corsHandler.Handler)

	router.Use(
		handler.FileServerMiddleware,
		chiMiddleware.Logger,
		chiMiddleware.Recoverer,
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
