package app

import (
	"github/madi-api/app/core"
	"github/madi-api/app/db"
	"github/madi-api/app/handler"
	"github/madi-api/config"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"

	log "github.com/sirupsen/logrus"
)

type App struct {
	Router *mux.Router
	DB     *mongo.Database
}

func (app *App) Run(host string) {
	handler := app.setupCors().Handler(app.Router)
	log.Printf("Server is listening on http://%s\n", host)

	err := http.ListenAndServe(host, requestLogger(handler))
	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) setupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", ""},
	})

}

func (app *App) setRouters() {
	app.Post("/signup-with-email", app.handleRequest(handler.SignUpAccountWithEmail))
	app.Post("/signin-with-email", app.handleRequest(handler.SignInAccountWithEmail))
	app.Post("/signin-with-social", app.handleRequest(handler.SignInAccountWithSocial))

}

func (app *App) UseMiddleware(middleware mux.MiddlewareFunc) {
	app.Router.Use(middleware)
	app.Router = app.Router.PathPrefix(core.PrefixAPI).Subrouter()

}

func (app *App) createIndexes() {
	account := app.DB.Collection(core.AccountCollection)

	keys := bsonx.Doc{
		{Key: "email", Value: bsonx.Int32(1)},
		{Key: "social.socialId", Value: bsonx.Int32(1)},
	}
	err := db.SetIndexes(account, keys)

	if err != nil {
		log.Fatal(err)
	}
	keys = bsonx.Doc{
		{Key: "email", Value: bsonx.Int32(1)},
	}
	db.SetIndexes(account, keys)
}

func (app *App) Initialize(config *config.Config) {
	app.DB = db.InitialConnection(config.MongoDB(), config.MongoURI())
	app.createIndexes()

	app.Router = mux.NewRouter()
	app.UseMiddleware(handler.JSONContentTypeMiddleware)
	app.setRouters()
}

func requestLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"time":    time.Now().Format(time.RFC822),
			"method":  r.Method,
			"url":     r.URL,
			"agent":   r.UserAgent(),
			"request": "Router",
		}).Info("A request has been received")
		handler.ServeHTTP(w, r)
	})
}

func (app *App) handleRequest(handler func(db *mongo.Database, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(app.DB, w, r)
	}
}

func (app *App) Get(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("GET").Queries(queries...)
}

func (app *App) Post(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("POST").Queries(queries...)
}

func (app *App) Put(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("PUT").Queries(queries...)
}

func (app *App) Patch(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("PATCH").Queries(queries...)
}

func (app *App) Delete(path string, endpoint http.HandlerFunc, queries ...string) {
	app.Router.HandleFunc(path, endpoint).Methods("DELETE").Queries(queries...)
}
