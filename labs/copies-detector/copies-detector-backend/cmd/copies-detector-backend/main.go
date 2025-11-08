package main

import (
	"log/slog"
	"os"

	"github.com/fatalistix/networks-and-telecommunications/copies-detector-backend/internal/app"
	"github.com/fatalistix/networks-and-telecommunications/copies-detector-backend/internal/util"
	"github.com/fatalistix/slogattr"
	"github.com/golang-cz/devslog"
)

func main() {
	log := setupLogger()

	a := app.New(log, func(port int) {
		err := util.WriteBigEndianUint16(os.Stdout, uint16(port))
		if err != nil {
			log.Error("Error writing bytes", slogattr.Err(err))
		}

		err = os.Stdout.Close()
		if err != nil {
			log.Error("Error closing stdout", slogattr.Err(err))
		}
	})

	log.Info("Starting application")

	go func() {
		if err := a.Run(); err != nil {
			log.Error("Error starting app", slogattr.Err(err))
		}
	}()

	_, err := os.Stdin.Read(make([]byte, 1))
	if err != nil {
		log.Warn("Error reading stdin, parent may be dead", slogattr.Err(err))
	}

	log.Info("Shutting down application")

	if err := a.Stop(); err != nil {
		log.Error("Error stopping app", slogattr.Err(err))
	}

	log.Info("Application stopped")
}

func setupLogger() *slog.Logger {
	return slog.New(
		devslog.NewHandler(os.Stderr, nil),
	)
}
