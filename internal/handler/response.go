package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		logrus.Errorf("error writing json to the stream: %+w\n", err)
	}
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	respondWithJSON(w, code, errorResponse{Message: err.Error()})
}
