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

	"github.com/joffrua/go-famtree/config"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	_ "github.com/joffrua/go-famtree/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type builder struct {
	cfg    *config.Config
	router *mux.Router
	server *http.Server
}

func NewBuilder(cfg *config.Config) *builder {
	b := new(builder)
	b.cfg = cfg
	b.router = mux.NewRouter()
	b.router.Use(Logging)

	b.server = &http.Server{
		Addr:    net.JoinHostPort("", b.cfg.HTTP.Port),
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

func (b *builder) AddSwagger(path string) {
	b.router.PathPrefix(path).Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
}

func (b *builder) Start() {
	go func() {
		if err := b.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("could not listen on port %s: %+v\n", b.cfg.HTTP.Port, err)
		}
	}()
	log.Infof("http server started on port: %s", b.cfg.HTTP.Port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	sig := <-c
	log.Infof("got '%s' signal, http server is shutting down...", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := b.server.Shutdown(ctx); err != nil {
		log.Panicf("could not gracefully shutdown http server: %+v", err)
	}
	log.Info("http server gracefully stopped")
}
