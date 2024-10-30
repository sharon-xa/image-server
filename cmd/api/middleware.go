package main

import (
	"log/slog"
	"net/http"
)

func (app *Application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ip     = r.RemoteAddr
			proto  = r.Proto
			method = r.Method
			uri    = r.RequestURI
		)

		app.logger.Info(
			"received request",
			slog.Any("ip", ip),
			slog.Any("proto", proto),
			slog.Any("method", method),
			slog.Any("uri", uri),
		)

		next.ServeHTTP(w, r)
	})
}
