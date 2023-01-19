package handler

import (
	"encoding/json"
	"github/madi-api/app/model"
	"net/http"
)

func ResponseWriter(res http.ResponseWriter, statusCode int, message string, data interface{}) error {
	res.WriteHeader(statusCode)
	httpResponse := model.NewResponse(statusCode, message, data)
	err := json.NewEncoder(res).Encode(httpResponse)
	return err
}
