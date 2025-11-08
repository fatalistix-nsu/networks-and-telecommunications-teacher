package copies_detectors

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/fatalistix/slogattr"
)

type GetResponse struct {
	CopiesDetectorsIds []string `json:"copies_detectors_ids"`
}

type GetService interface {
	GetAllIds() []string
}

func MakeGetHandler(log *slog.Logger, service GetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ids := service.GetAllIds()

		resp := GetResponse{
			CopiesDetectorsIds: ids,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Error("Failed to write response body", slogattr.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte("Failed to write response body"))
			if err != nil {
				log.Error("Failed to write response body", slogattr.Err(err))
			}
			return
		}
	}
}
