package id

import (
	"log/slog"
	"net/http"

	"github.com/fatalistix/networks-and-telecommunications/copies-detector-backend/internal/http/middleware"
)

func RegisterHandlers(
	log *slog.Logger,
	mux *http.ServeMux,
	deleteService DeleteService,
	getService GetService,
) {
	mux.HandleFunc("DELETE /copies_detectors/{id}", middleware.WithCors(middleware.WithLogging(log, MakeDeleteHandler(log, deleteService))))
	mux.HandleFunc("OPTIONS /copies_detectors/{id}", middleware.WithCors(middleware.WithLogging(log, MakeOptionsHandler())))
	mux.HandleFunc("GET /copies_detectors/{id}", middleware.WithCors(middleware.WithLogging(log, MakeGetHandler(log, getService))))
}
