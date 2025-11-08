package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	copiesdetectors "github.com/fatalistix/networks-and-telecommunications/copies-detector-backend/internal/http/handler/copies_detectors"
	copiesdetectorsid "github.com/fatalistix/networks-and-telecommunications/copies-detector-backend/internal/http/handler/copies_detectors/id"
	"github.com/fatalistix/networks-and-telecommunications/copies-detector-backend/internal/service"
	"github.com/fatalistix/networks-and-telecommunications/copies-detector-backend/internal/util"
)

type ConnectListener = func(port int)

type Server struct {
	server          *http.Server
	connectListener ConnectListener
}

func NewServer(log *slog.Logger, connectListener ConnectListener, s *service.CopiesDetectorService) *Server {
	mux := http.NewServeMux()

	copiesdetectors.RegisterHandlers(log, mux, s, s)
	copiesdetectorsid.RegisterHandlers(log, mux, s, s)

	return &Server{
		server: &http.Server{
			Handler: mux,
		},
		connectListener: connectListener,
	}
}

func (s *Server) Run() error {
	l, port, err := util.ListenTCPAnyPort()
	if err != nil {
		return fmt.Errorf("failed to listen any tcp port: %w", err)
	}

	go s.connectListener(port)

	if err := s.server.Serve(l); err != nil {
		return fmt.Errorf("could not start http server: %w", err)
	}

	return nil
}

func (s *Server) Shutdown() error {
	if err := s.server.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("could not shutdown http server: %w", err)
	}

	return nil
}
