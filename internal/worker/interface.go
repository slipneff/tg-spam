package worker

import (
	"context"

	"github.com/gotd/td/telegram"
	"github.com/slipneff/tg-spam/internal/utils/config"
)

type Worker struct {
	Client  *telegram.Client
	storage Storage
	cfg     *config.Config
}

type Storage interface {
	GetLastMessageID(ctx context.Context, key string) (int, error)
	SetLastMessageID(ctx context.Context, key string, value int) error

	GetSessions(ctx context.Context, n int) ([]string, error)
}

func NewWorker(Client *telegram.Client, cfg *config.Config) *Worker {
	return &Worker{Client: Client, cfg: cfg}
}
