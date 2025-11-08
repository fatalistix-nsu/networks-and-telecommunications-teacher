package service

import (
	"log/slog"
	"slices"
	"sync"
	"time"

	"github.com/fatalistix/networks-and-telecommunications/copies-detector/internal/constant"
	"github.com/fatalistix/networks-and-telecommunications/copies-detector/pkg/model"
)

type CopiesService struct {
	log          *slog.Logger
	mutex        *sync.RWMutex
	activeCopies map[string]time.Time
}

func NewCopiesService(log *slog.Logger) *CopiesService {
	return &CopiesService{
		log:          log,
		mutex:        &sync.RWMutex{},
		activeCopies: make(map[string]time.Time),
	}
}

func (s *CopiesService) AddOrRefresh(id string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.activeCopies[id] = time.Now()
}

func (s *CopiesService) Remove(id string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	delete(s.activeCopies, id)
}

func (s *CopiesService) CleanTimedOut() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	now := time.Now()

	toRemove := make([]string, 0, len(s.activeCopies))

	for k, v := range s.activeCopies {
		diff := now.Sub(v)
		if diff > constant.RemoveTimeout {
			toRemove = append(toRemove, k)
		}
	}

	for _, id := range toRemove {
		delete(s.activeCopies, id)
	}
}

func (s *CopiesService) GetActiveCopies() []model.ActiveCopyWithLastRefresh {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	sortedActiveCopies := make([]string, 0, len(s.activeCopies))

	for k := range s.activeCopies {
		sortedActiveCopies = append(sortedActiveCopies, k)
	}

	slices.Sort(sortedActiveCopies)

	result := make([]model.ActiveCopyWithLastRefresh, 0, len(sortedActiveCopies))

	for _, v := range sortedActiveCopies {
		result = append(result, model.ActiveCopyWithLastRefresh{
			ActiveCopy:  v,
			LastRefresh: s.activeCopies[v],
		})
	}

	return result
}
