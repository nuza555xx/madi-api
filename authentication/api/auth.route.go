package auth

import (
	"github.com/gorilla/mux"
)

func Routes(routers *mux.Router) *mux.Router {
	router := routers.PathPrefix("/auth").Subrouter()
	router.HandleFunc("/register-with-email", registerWithEmail()).Methods("POST")
	// router.HandleFunc("/login-with-email", loginWithEmail()).Methods("POST")

	return router
}