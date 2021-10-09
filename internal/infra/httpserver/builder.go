package httpserver

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/joffrua/go-famtree/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type builder struct {
	router *mux.Router
	port   string
}

func NewBuilder() *builder {
	b := new(builder)
	b.router = mux.NewRouter()

	b.port = os.Getenv("PORT")
	if b.port == "" {
		b.port = "8080"
	}

	return b
}

func (b *builder) AddRoute(method, path string, handlerFunc http.HandlerFunc) {
	b.router.Methods(method).Path(path).HandlerFunc(handlerFunc)
}

func (b *builder) ServeStatic(path, dir string) {
	b.router.PathPrefix(path).Handler(http.FileServer(http.Dir(dir)))
}

func (b *builder) AddSwagger(path string) {
	b.router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
}

func (b *builder) ListenAndServe() {
	fmt.Println("Start server on: ", b.port)
	if err := http.ListenAndServe(":"+b.port, b.router); err != nil {
		panic(err.Error())
	}
}
