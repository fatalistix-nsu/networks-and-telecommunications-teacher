package copies_detector

import (
	"errors"
	"log/slog"
	"sync"

	"github.com/fatalistix/networks-and-telecommunications/copies-detector/internal/constant"
	"github.com/fatalistix/networks-and-telecommunications/copies-detector/internal/multicast"
	"github.com/fatalistix/networks-and-telecommunications/copies-detector/internal/service"
	"github.com/fatalistix/networks-and-telecommunications/copies-detector/internal/worker"
	"github.com/fatalistix/networks-and-telecommunications/copies-detector/pkg/model"
	"github.com/fatalistix/slogattr"
)

type CopiesDetector struct {
	log      *slog.Logger
	s        *service.CopiesService
	sender   *multicast.Sender
	listener *multicast.Listener
	cleaner  *worker.Cleaner
}

func New(log *slog.Logger, name, host string, port int) *CopiesDetector {
	s := service.NewCopiesService(log)

	joinOrRefreshHandler := multicast.NewJoinOrRefreshHandler(log, s)
	leaveHandler := multicast.NewLeaveHandler(log, s)

	listener := multicast.NewListener(log, host, port)
	listener.SetHandler(multicast.JoinOrRefresh, joinOrRefreshHandler)
	listener.SetHandler(multicast.Leave, leaveHandler)

	sender := multicast.NewSender(log, host, port, name, constant.RefreshDuration)
	cleaner := worker.NewCleaner(log, s)

	return &CopiesDetector{
		log:      log,
		s:        s,
		sender:   sender,
		listener: listener,
		cleaner:  cleaner,
	}
}

func (a *CopiesDetector) Run() error {
	wg := &sync.WaitGroup{}
	errChan := make(chan error)

	wg.Go(func() {
		if err := a.sender.Run(); err != nil {
			a.log.Error("Error running sender", slogattr.Err(err))
			errChan <- err
		}
	})

	wg.Go(func() {
		if err := a.listener.Run(); err != nil {
			a.log.Error("Error running listener", slogattr.Err(err))
			errChan <- err
		}
	})

	wg.Go(func() {
		a.cleaner.Run()
	})

	go func() {
		wg.Wait()
		close(errChan)
	}()

	errs := make([]error, 0)
	for err := range errChan {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func (a *CopiesDetector) GetActiveCopies() []model.ActiveCopyWithLastRefresh {
	return a.s.GetActiveCopies()
}

func (a *CopiesDetector) Stop() {
	a.sender.Shutdown()
	a.listener.Shutdown()
	a.cleaner.Shutdown()
}
