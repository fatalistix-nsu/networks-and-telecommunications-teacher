package copies_detectors

import (
	"log/slog"
	"net/http"

	"github.com/fatalistix/networks-and-telecommunications/copies-detector-backend/internal/http/middleware"
)

func RegisterHandlers(
	log *slog.Logger,
	mux *http.ServeMux,
	getService GetService,
	postService PostService,
) {
	mux.HandleFunc("GET /copies_detectors", middleware.WithCors(middleware.WithLogging(log, MakeGetHandler(log, getService))))
	mux.HandleFunc("OPTIONS /copies_detectors", middleware.WithCors(middleware.WithLogging(log, MakeOptionsHandler())))
	mux.HandleFunc("POST /copies_detectors", middleware.WithCors(middleware.WithLogging(log, MakePostHandler(log, postService))))
}
