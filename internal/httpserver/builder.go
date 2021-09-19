package httpserver

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type builder struct {
	router *mux.Router
}

func NewBuilder() *builder {
	b := new(builder)
	b.router = mux.NewRouter()
	return b
}

func (b *builder) AddRoute(method, path string, handlerFunc http.HandlerFunc) {
	b.router.Methods(method).Path(path).HandlerFunc(handlerFunc)
}

func (b *builder) ServeStatic(path, dir string) {
	b.router.PathPrefix(path).Handler(http.FileServer(http.Dir(dir)))
}

func (b *builder) ListenAndServe() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(":"+port, b.router); err != nil {
		panic(err.Error())
	}
}
