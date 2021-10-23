package httpserver

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		log.WithFields(log.Fields{
			"uri":      r.RequestURI,
			"method":   r.Method,
			"duration": time.Since(start),
		}).Info()
	})
}
