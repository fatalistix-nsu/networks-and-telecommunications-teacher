package worker

import (
	"log/slog"
	"sync"
	"time"
)

const (
	checkDelay = time.Second
)

type CleanerService interface {
	CleanTimedOut()
}

type Cleaner struct {
	log      *slog.Logger
	service  CleanerService
	stopChan chan bool
	stopWg   *sync.WaitGroup
}

func NewCleaner(log *slog.Logger, service CleanerService) *Cleaner {
	stopWg := &sync.WaitGroup{}
	stopWg.Add(1)

	return &Cleaner{
		log:      log,
		service:  service,
		stopChan: make(chan bool),
		stopWg:   stopWg,
	}
}

func (w *Cleaner) Run() {
	defer w.stopWg.Done()

	t := time.NewTicker(checkDelay)
	defer t.Stop()

	for {
		select {
		case <-w.stopChan:
			w.log.Info("Stopping cleaning worker")
			return
		case <-t.C:
			w.service.CleanTimedOut()
		}
	}
}

func (w *Cleaner) Shutdown() {
	w.stopChan <- true
	w.stopWg.Wait()
}
