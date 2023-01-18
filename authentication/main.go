package main

import (
	"log"
	auth "madi-api/api"
	"madi-api/common"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	port := common.DotEnvVariable("PORT")

	common.DebugLogger("Application running on :: " + port)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	router := mux.NewRouter()

	auth.Routes(router)

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	handler := c.Handler(router)

	http.ListenAndServe(":"+port, common.RequestLogger(handler))
}
