package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// App export
type App struct {
	Router *mux.Router
}

func (app *App) initialiseRoutes() {
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/", get_all)
}

func (app *App) run() {
	log.Fatal(http.ListenAndServe(":8000", app.Router))
}
