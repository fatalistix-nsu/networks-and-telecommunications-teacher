package app

import (
	"fmt"
	"log/slog"

	"github.com/fatalistix/networks-and-telecommunications/copies-detector-backend/internal/http/server"
	"github.com/fatalistix/networks-and-telecommunications/copies-detector-backend/internal/service"
)

type ConnectListener = func(port int)

type App struct {
	log        *slog.Logger
	httpServer *server.Server
}

func New(
	log *slog.Logger,
	connectListener ConnectListener,
) *App {
	s := service.NewCopiesDetectorService(log)

	httpServer := server.NewServer(log, connectListener, s)

	return &App{
		log:        log,
		httpServer: httpServer,
	}
}

func (a *App) Run() error {
	if err := a.httpServer.Run(); err != nil {
		return fmt.Errorf("run http server: %w", err)
	}

	return nil
}

func (a *App) Stop() error {
	if err := a.httpServer.Shutdown(); err != nil {
		return fmt.Errorf("could not shutdown http server: %w", err)
	}

	return nil
}
