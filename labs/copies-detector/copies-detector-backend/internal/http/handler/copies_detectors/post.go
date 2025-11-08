package copies_detectors

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/fatalistix/slogattr"
)

type PostRequest struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Name string `json:"name"`
}

type PostResponse struct {
	Id string `json:"id"`
}

type PostService interface {
	RunCopiesDetector(host string, port int, name string) (string, error)
}

func MakePostHandler(log *slog.Logger, service PostService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bodyReader := r.Body
		if bodyReader == nil {
			log.Error("Request has no body")
			w.WriteHeader(http.StatusUnprocessableEntity)
			_, err := w.Write([]byte("empty body"))
			if err != nil {
				log.Error("Failed to write response body", slogattr.Err(err))
			}
			return
		}

		defer func() {
			err := bodyReader.Close()
			if err != nil {
				log.Error("Failed to close body", slogattr.Err(err))
			}
		}()

		req := PostRequest{}

		decoder := json.NewDecoder(bodyReader)
		err := decoder.Decode(&req)
		if err != nil {
			log.Error("Failed to decode request", slogattr.Err(err))
			w.WriteHeader(http.StatusUnprocessableEntity)
			_, err := w.Write([]byte("invalid request"))
			if err != nil {
				log.Error("Failed to write response body", slogattr.Err(err))
			}
			return
		}

		id, err := service.RunCopiesDetector(req.Host, req.Port, req.Name)
		if err != nil {
			log.Error("Failed to create copies_detectors", slogattr.Err(err))
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte("Failed to create copies_detectors"))
			if err != nil {
				log.Error("Failed to write response body", slogattr.Err(err))
			}
			return
		}

		resp := PostResponse{
			Id: id,
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(http.StatusCreated)
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
