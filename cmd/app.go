package app

import (
	"github.com/gorilla/mux"
	"moq.com/test/cmd/database"
)

type App struct {
	Router *mux.Router
	DB		database.RandomDb
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}
	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/api/v1/posts", a.CreateHandler()).Methods("POST")
	a.Router.HandleFunc("/api/v1/posts", a.GetHandler()).Methods("GET")
}
