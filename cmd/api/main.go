package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Application struct {
	logger     *slog.Logger
	imgDirPath string
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	imgDirPath := flag.String("path", "/home/sharon/images/", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &Application{
		logger:     logger,
		imgDirPath: *imgDirPath,
	}

	err := app.initDir(*imgDirPath)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	server := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Info("Starting server", slog.Any("PORT", server.Addr))
	err = server.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
