package httpserver

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

type ContextKey string

const ContextKeyRequestID ContextKey = "requestID"

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		buf := []byte("{}")
		if body != nil {

			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Error().Err(err).Msg("Cannot read out the request body")
				return
			}
			if len(b) > 0 {
				buf = b
			}
			r.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		}

		log.Trace().Str("method", r.Method).Str("url", r.URL.String()).Interface("headers", r.Header).RawJSON("request_body", buf).Msg("The incoming request")

		next.ServeHTTP(w, r)

		r.Body = body
	})
}
