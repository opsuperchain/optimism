package frontend

import (
	"context"
)

type SyncBackend interface {
	GetSession(ctx context.Context) error
	DeleteSession(ctx context.Context) error
	ListSessions(ctx context.Context) ([]string, error)
}

type SyncFrontend struct {
	SyncBackend
}

func NewSyncFrontend(b SyncBackend) *SyncFrontend {
	return &SyncFrontend{SyncBackend: b}
}
