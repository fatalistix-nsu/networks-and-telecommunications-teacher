package id

import (
	"log/slog"
	"net/http"

	"github.com/fatalistix/slogattr"
)

type DeleteService interface {
	StopCopiesDetector(id string) error
}

func MakeDeleteHandler(log *slog.Logger, service DeleteService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "" {
			log.Error("Missing id path variable")
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte("missing id path variable"))
			if err != nil {
				log.Error("Failed to write response body", slogattr.Err(err))
			}
			return
		}

		err := service.StopCopiesDetector(id)
		if err != nil {
			log.Error("Failed to create copies_detectors", slogattr.Err(err))
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte("Failed to create copies_detectors"))
			if err != nil {
				log.Error("Failed to write response body", slogattr.Err(err))
			}
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
