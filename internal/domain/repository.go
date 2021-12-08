package domain

import "context"

type HelloRepository interface {
	GetHello(ctx context.Context, user string) (*Hello, error)
}

type HistoryRepository interface {
	AddRecord(ctx context.Context, user, r *HistoryRecord) error
	GetHistory(ctx context.Context, user string) (*History, error)
}
