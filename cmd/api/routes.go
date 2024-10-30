package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *Application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /image", app.saveImage)
	mux.HandleFunc("GET /image", app.getImage)
	mux.HandleFunc("DELETE /image", app.deleteImage)

	loggedHandlers := alice.New(app.logRequest)

	return loggedHandlers.Then(mux)
}
