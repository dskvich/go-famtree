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

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	_ "github.com/joffrua/go-famtree/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type builder struct {
	router *mux.Router
	server *http.Server
	port   string
}

func NewBuilder() *builder {
	b := new(builder)
	b.router = mux.NewRouter()
	b.router.Use(Logging)

	b.port = os.Getenv("PORT")
	if b.port == "" {
		b.port = "8080"
	}

	log.Infof("get env PORT=%s", b.port)

	b.server = &http.Server{
		Addr:    net.JoinHostPort("localhost", b.port),
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
			log.Errorf("could not listen on port %s: %+v\n", b.port, err)
		}
	}()
	log.Infof("http server started on port: %s", b.port)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt
	log.Info("http server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := b.server.Shutdown(ctx); err != nil {
		log.Panicf("could not gracefully shutdown http server: %+v", err)
	}
	log.Info("http server gracefully stopped")
}
