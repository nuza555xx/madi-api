package handler

import (
	"encoding/json"
	"net/http"
)

func ResponseWithError(res http.ResponseWriter, statusCode int, message string) {
	ResponseWithJSON(res, statusCode, map[string]string{"error": message})
}

func ResponseWithValidationError(res http.ResponseWriter, statusCode int, errors map[string][]string) {
	ResponseWithJSON(res, statusCode, map[string]interface{}{"error": "Invalid input", "errors": errors})
}

func ResponseWithJSON(res http.ResponseWriter, statusCode int, payload interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	json.NewEncoder(res).Encode(payload)
}
