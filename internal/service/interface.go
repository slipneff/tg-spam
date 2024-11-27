package service

import (
	"context"

	"github.com/slipneff/tg-spam/internal/models"
	"github.com/slipneff/tg-spam/internal/utils/config"
	"github.com/slipneff/tg-spam/internal/worker"
)

type Service struct {
	Worker  *worker.Worker
	storage Storage
	cfg     *config.Config
}

type Storage interface {
	GetChannels(ctx context.Context) ([]*models.Channel, error)
}

func NewService(Worker *worker.Worker, cfg *config.Config) *Service {
	return &Service{Worker: Worker, cfg: cfg}
}
