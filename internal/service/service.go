package service

import (
	"context"

	"github.com/slipneff/tg-spam/internal/models"
)

func (s *Service) getChannels(ctx context.Context) ([]*models.Channel, error) {
	return s.storage.GetChannels(ctx)
}

func (s *Service) CatchingChannels(ctx context.Context) error {
	channels, err := s.getChannels(ctx)
	if err != nil {
		return err
	}

	for _, channel := range channels {
		for {
			select {
			case <-ctx.Done():
				return nil
			default:
				go s.Worker.CatchLastPost(ctx, channel.Name)
			}
		}
	}
	return nil
}

