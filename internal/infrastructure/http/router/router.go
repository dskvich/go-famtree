package router

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

type Config struct {
	AllowedOrigins []string
	AllowedHeaders []string
	Logger         *logrus.Entry
}

func (cfg *Config) validate() error {
	var err error
	if cfg.Logger == nil {
		cfg.Logger = logrus.NewEntry(&logrus.Logger{Out: ioutil.Discard})
	}
	return err
}

type router struct {
	cfg Config
	mux *mux.Router
}

func New(cfg Config) *router {
	return &router{
		mux: mux.NewRouter().StrictSlash(true),
	}
}

func (r *router) Subrouter(pathPrefix string) *router {
	return &router{
		mux: r.mux.PathPrefix(pathPrefix).Subrouter(),
	}
}

func (r *router) Add(method, path string, handlerFunc http.HandlerFunc) *mux.Route {
	handler := r.corsHandler(r.requestLogger(handlerFunc))
	return r.add(method, path, handler)
}

func (r *router) add(method, path string, handler http.Handler) *mux.Route {
	return r.mux.Methods(method).Path(path).Handler(handler)
}

func (r *router) requestLogger(inner http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, req)
		r.cfg.Logger.Debugf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})

}

func (r *router) corsHandler(handler http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   r.cfg.AllowedOrigins,
		AllowedHeaders:   r.cfg.AllowedHeaders,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
	}).Handler(handler)
}

func (r *router) Handler() *mux.Router {
	return r.mux
}
