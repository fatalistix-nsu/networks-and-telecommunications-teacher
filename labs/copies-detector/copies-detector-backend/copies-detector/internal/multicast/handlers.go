package multicast

import "log/slog"

type JoinOrRefreshService interface {
	AddOrRefresh(id string)
}

func buildId(addr, name string) string {
	return name + "@" + addr
}

func NewJoinOrRefreshHandler(log *slog.Logger, s JoinOrRefreshService) Handler {
	return func(addr string, name string) {
		log.Info("Received join", slog.String("from", addr), slog.String("name", name))

		id := buildId(addr, name)
		s.AddOrRefresh(id)
	}
}

type LeaveService interface {
	Remove(id string)
}

func NewLeaveHandler(log *slog.Logger, s LeaveService) Handler {
	return func(addr string, name string) {
		log.Info("Received leave", slog.String("from", addr), slog.String("name", name))

		id := buildId(addr, name)
		s.Remove(id)
	}
}
