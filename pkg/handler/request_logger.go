package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/log"
)

func RequestLogger(next http.Handler) http.Handler {
	return middleware.RequestLogger(&logFormatter{})(next)
}

type logFormatter struct{}

func (c *logFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	request := map[string]interface{}{}

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		request["id"] = reqID
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	request["method"] = r.Method
	request["remote"] = r.RemoteAddr
	request["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

	request["params"] = serializeParams(r)

	fields := map[string]interface{}{}
	fields["request"] = request

	return &CustomLogEntry{
		fields: fields,
	}
}

func serializeParams(r *http.Request) string {
	var params []string

	// Serialize query parameters
	queryParams := r.URL.Query()
	for key, values := range queryParams {
		for _, value := range values {
			params = append(params, fmt.Sprintf("%s=%s", key, value))
		}
	}

	// Serialize request body
	if r.Body != nil {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error().Err(err).Msg("reading request body")
		}

		r.Body = io.NopCloser(bytes.NewBuffer(body))

		if len(body) > 0 {
			data := make(map[string]interface{})
			err = json.Unmarshal(body, &data)
			if err != nil {
				log.Error().Err(err).Msg("parsing request payload")
			} else {
				for key, value := range data {
					params = append(params, fmt.Sprintf("%s=%v", key, value))
				}
			}
		}
	}

	return strings.Join(params, ", ")
}

type CustomLogEntry struct {
	fields map[string]interface{}
	errors []map[string]interface{}
}

func (c *CustomLogEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	response := map[string]interface{}{}
	response["status"] = status
	response["bytes"] = bytes
	response["elapsed"] = float64(elapsed.Nanoseconds()) / 1000000.0

	c.fields["response"] = response

	if len(c.errors) > 0 {
		c.fields["errors"] = c.errors
		log.Error().Fields(c.fields).Msgf("request failed (%d)", status)
	} else {
		log.Debug().Fields(c.fields).Msgf("request complete (%d)", status)
	}
}

func (c *CustomLogEntry) Panic(v interface{}, stack []byte) {
	err := map[string]interface{}{}
	err["message"] = fmt.Sprintf("%+v", v)
	err["stack"] = string(stack)

	c.errors = append(c.errors, err)
}
