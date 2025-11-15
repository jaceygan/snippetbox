package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger *slog.Logger
}

func main() {
	port := flag.String("port", "4000", "Port to run the web server on")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo, AddSource: true})) // &var returns memory address of var
	// logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app := &application{logger: logger}

	db, err := openDB(*dsn)

	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	logger.Info("Starting server", slog.Any("port", *port)) // *port dereferences the pointer to get the actual value
	err = http.ListenAndServe(":"+*port, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
