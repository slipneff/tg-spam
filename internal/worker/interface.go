package worker

import (
	"context"

	"github.com/gotd/td/telegram"
	"github.com/slipneff/tg-spam/internal/models"
	"github.com/slipneff/tg-spam/internal/utils/config"
	"github.com/slipneff/tg-spam/pkg/gpt"
)

type Worker struct {
	Client    *telegram.Client
	storage   Storage
	cfg       *config.Config
	gptClient *gpt.Client
}

type Storage interface {
	GetLastMessageID(ctx context.Context, key string) (int, error)
	SetLastMessageID(ctx context.Context, key string, value int) error

	GetSessions(ctx context.Context, n int) ([]string, error)
	GetSessionById(ctx context.Context, id string) (*models.Session, error)
}

func NewWorker(Client *telegram.Client, gptClient *gpt.Client, cfg *config.Config) *Worker {
	return &Worker{Client: Client, gptClient: gptClient, cfg: cfg}
}
