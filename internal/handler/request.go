package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func decodeFromJSON(body io.Reader, v any) (int, error) {
	//r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	d := json.NewDecoder(body)
	d.DisallowUnknownFields()

	err := d.Decode(&v)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var maxBytesError *http.MaxBytesError

		switch {
		case errors.As(err, &syntaxError):
			return http.StatusBadRequest, fmt.Errorf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return http.StatusBadRequest, fmt.Errorf("Request body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return http.StatusBadRequest, fmt.Errorf("Request body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return http.StatusBadRequest, fmt.Errorf("Request body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return http.StatusBadRequest, fmt.Errorf("Request body contains unknown field %s", fieldName)
		case errors.Is(err, io.EOF):
			return http.StatusBadRequest, fmt.Errorf("Request body must not be empty")
		case errors.As(err, &maxBytesError):
			return http.StatusRequestEntityTooLarge, fmt.Errorf("Request body must not be larger than %d bytes", maxBytesError.Limit)
		default:
			return http.StatusInternalServerError, fmt.Errorf("Internal Server Error")
		}
	}

	if err = d.Decode(&struct{}{}); err != io.EOF {
		return http.StatusBadRequest, fmt.Errorf("Request body must only contain a single JSON object")
	}

	return http.StatusOK, nil
}
