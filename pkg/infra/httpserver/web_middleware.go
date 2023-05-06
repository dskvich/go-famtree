package httpserver

import (
	"net/http"
	"strings"
)

func FileServerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			next.ServeHTTP(w, r)
		} else {
			http.FileServer(http.Dir("./build")).ServeHTTP(w, r)
		}
	})
}
