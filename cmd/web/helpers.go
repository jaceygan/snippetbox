package main

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)

type application struct {
	logger *slog.Logger
}

func (app *application) serverInfo(r *http.Request, message string) {
	var (
		method = r.Method
		url    = r.URL.RequestURI()
	)
	app.logger.Info(message, slog.String("method", method), slog.String("url", url))
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		url    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)
	app.logger.Error(err.Error(), slog.String("method", method), slog.String("url", url), slog.String("trace", trace))
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
