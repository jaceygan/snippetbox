package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	port := flag.String("port", "4000", "Port to run the web server on")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo, AddSource: true})) // &var returns memory address of var
	// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := &application{logger: logger}

	logger.Info("Starting server", slog.Any("port", *port)) // *port dereferences the pointer to get the actual value
	err := http.ListenAndServe(":"+*port, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
