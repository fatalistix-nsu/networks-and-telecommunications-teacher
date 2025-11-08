package service

import (
	"fmt"
	"log/slog"
	"sync"

	copiesdetector "github.com/fatalistix/networks-and-telecommunications/copies-detector/pkg"
	"github.com/fatalistix/networks-and-telecommunications/copies-detector/pkg/model"
	"github.com/fatalistix/slogattr"
	"github.com/google/uuid"
)

type runningCopiesDetector struct {
	cd *copiesdetector.CopiesDetector
}

type CopiesDetectorService struct {
	log    *slog.Logger
	mutex  *sync.Mutex
	copies map[string]*runningCopiesDetector
}

func NewCopiesDetectorService(log *slog.Logger) *CopiesDetectorService {
	return &CopiesDetectorService{
		log:    log,
		mutex:  &sync.Mutex{},
		copies: make(map[string]*runningCopiesDetector),
	}
}

func (s *CopiesDetectorService) RunCopiesDetector(host string, port int, name string) (string, error) {
	id := uuid.New().String()
	cd := copiesdetector.New(s.log, name, host, port)

	lcd := &runningCopiesDetector{
		cd: cd,
	}

	go func() {
		if err := lcd.cd.Run(); err != nil {
			s.log.Error(
				"Failed to run copies detector",
				slog.String("host", host),
				slog.String("name", name),
				slog.Int("port", port),
				slogattr.Err(err),
			)
		}
	}()

	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.copies[id] = lcd

	return id, nil
}

func (s *CopiesDetectorService) GetAllIds() []string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	ids := make([]string, 0, len(s.copies))
	for id := range s.copies {
		ids = append(ids, id)
	}

	return ids
}

func (s *CopiesDetectorService) GetActiveCopies(id string) ([]model.ActiveCopyWithLastRefresh, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	lcd, ok := s.copies[id]
	if !ok {
		return nil, fmt.Errorf("copies_detectors %s not found", id)
	}

	return lcd.cd.GetActiveCopies(), nil
}

func (s *CopiesDetectorService) StopCopiesDetector(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	lcd, ok := s.copies[id]
	if !ok {
		return fmt.Errorf("copies_detectors %s not found", id)
	}

	delete(s.copies, id)

	lcd.cd.Stop()

	return nil
}
