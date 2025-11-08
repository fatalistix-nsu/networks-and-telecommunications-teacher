package id

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/fatalistix/networks-and-telecommunications/copies-detector/pkg/model"
	"github.com/fatalistix/slogattr"
)

type ActiveCopyWithLastRefreshDto struct {
	ActiveCopy  string    `json:"active_copy"`
	LastRefresh time.Time `json:"last_refresh"`
}

func mapActiveCopyWithLastRefreshModelToDto(m model.ActiveCopyWithLastRefresh) ActiveCopyWithLastRefreshDto {
	return ActiveCopyWithLastRefreshDto{
		ActiveCopy:  m.ActiveCopy,
		LastRefresh: m.LastRefresh,
	}
}

func mapActiveCopyWithLastRefreshModelListToDto(m []model.ActiveCopyWithLastRefresh) []ActiveCopyWithLastRefreshDto {
	result := make([]ActiveCopyWithLastRefreshDto, len(m))
	for i, v := range m {
		result[i] = mapActiveCopyWithLastRefreshModelToDto(v)
	}

	return result
}

type GetResponse struct {
	Id           string                         `json:"id"`
	ActiveCopies []ActiveCopyWithLastRefreshDto `json:"active_copies"`
}

type GetService interface {
	GetActiveCopies(id string) ([]model.ActiveCopyWithLastRefresh, error)
}

func MakeGetHandler(log *slog.Logger, service GetService) http.HandlerFunc {
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

		acm, err := service.GetActiveCopies(id)
		if err != nil {
			log.Error("Failed to get active copies", slogattr.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte("failed to get active copies"))
			if err != nil {
				log.Error("Failed to write response body", slogattr.Err(err))
			}
			return
		}

		acDto := mapActiveCopyWithLastRefreshModelListToDto(acm)

		resp := GetResponse{
			Id:           id,
			ActiveCopies: acDto,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
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
