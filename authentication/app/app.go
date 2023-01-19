package app

import (
	"github/madi-api/app/db"
	"github/madi-api/app/handler"
	"github/madi-api/config"
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type App struct {
	Router *mux.Router
	DB     *mongo.Database
}

func (app *App) Initialize(config *config.Config) {
	app.DB = db.InitialConnection("dev", config.MongoURI())
	app.createIndexes()

	app.Router = mux.NewRouter()
	app.UseMiddleware(handler.JSONContentTypeMiddleware)
	app.setRouters()
}

func (app *App) setRouters() {
	app.Post("/register-with-email", app.handleRequest(handler.CreateAccountWitEmail))
}

func (app *App) UseMiddleware(middleware mux.MiddlewareFunc) {
	app.Router.Use(middleware)
	app.Router = app.Router.PathPrefix("/auth").Subrouter()

}

func (app *App) createIndexes() {
	keys := bsonx.Doc{
		{Key: "email", Value: bsonx.Int32(1)},
	}
	account := app.DB.Collection("account")
	db.SetIndexes(account, keys)
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

func (app *App) Run(host string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	handler := c.Handler(app.Router)

	log.Printf("Server is listening on http://%s\n", host)

	err := http.ListenAndServe(host, requestLogger(handler))
	if err != nil {
		log.Fatal(err)
	}

}

func requestLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		color.Cyan("%s - [ Router ] 🥹  %s %s %s %s\n", time.Now().Format(time.RFC822), r.Host, r.Method, r.URL, r.UserAgent())
		handler.ServeHTTP(w, r)
	})
}

func (app *App) handleRequest(handler func(db *mongo.Database, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(app.DB, w, r)
	}
}
