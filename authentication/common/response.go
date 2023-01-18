package common

import (
	"encoding/json"
	"net/http"
)

type ErrorValidation struct {
	StatusCode int                 `json:"statusCode"`
	Message    string              `json:"message"`
	Errors     map[string][]string `json:"errors"`
}

type Error struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func ErrorResponse(m string, w http.ResponseWriter) {
	ErrorLogger(m)
	temp := Error{StatusCode: http.StatusBadRequest, Message: m}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(temp)
}

func SuccessResponse(fields interface{}, w http.ResponseWriter) {
	_, err := json.Marshal(fields)

	if err != nil {
		ErrorResponse(http.StatusText(500), w)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fields)
}

func ValidationResponse(fields map[string][]string, w http.ResponseWriter) {

	temp := ErrorValidation{StatusCode: http.StatusBadRequest, Message: http.StatusText(400), Errors: fields}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(temp)
}
