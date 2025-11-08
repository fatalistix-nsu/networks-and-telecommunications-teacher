package model

import "time"

type ActiveCopyWithLastRefresh struct {
	ActiveCopy  string
	LastRefresh time.Time
}
