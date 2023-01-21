package handler

import (
	"encoding/json"
	"github/madi-api/app/model"
	"log"
	"net/http"
)

func ResponseWriter(res http.ResponseWriter, statusCode int, message string, data interface{}) {
	res.WriteHeader(statusCode)
	httpResponse := model.NewResponse(statusCode, message, data)
	err := json.NewEncoder(res).Encode(httpResponse)
	if err != nil {
		log.Fatal(err)
	}
}
