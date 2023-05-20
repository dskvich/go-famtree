package handler

import (
	"net/http"
	"strings"
)

func FileServer(next http.Handler) http.Handler {
	fileServer := http.FileServer(http.Dir("./build"))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			next.ServeHTTP(w, r)
		} else {
			fileServer.ServeHTTP(w, r)
		}
	})
}
