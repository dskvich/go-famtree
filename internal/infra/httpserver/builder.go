package httpserver

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"github.com/joffrua/go-famtree/config"
)

type builder struct {
	cfg    *config.HTTP
	router *mux.Router
	server *http.Server
}

func NewBuilder(cfg *config.HTTP) *builder {
	b := new(builder)
	b.cfg = cfg
	b.router = mux.NewRouter().StrictSlash(true)
	b.router.Use(Logging)

	b.server = &http.Server{
		Addr:    net.JoinHostPort("", b.cfg.Port),
		Handler: b.router,
	}

	return b
}

func (b *builder) AddRoute(method, path string, handlerFunc http.HandlerFunc) {
	b.router.Methods(method).Path(path).HandlerFunc(handlerFunc)
}

func (b *builder) AddStaticDir(path, dir string) {
	b.router.PathPrefix(path).Handler(http.FileServer(http.Dir(dir)))
}

func (b *builder) Start() {
	go func() {
		if err := b.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Err(err).Msgf("Could not listen on port %s", b.cfg.Port)
		}
	}()
	log.Info().Msgf("HTTP server started on port: %s", b.cfg.Port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	sig := <-c
	log.Info().Msgf("Got '%s' signal, HTTP server is shutting down...", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := b.server.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Could not gracefully shutdown http server")
	}
	log.Info().Msg("HTTP server gracefully stopped")
}
