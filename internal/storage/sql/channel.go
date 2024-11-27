package sql

import (
	"context"

	"github.com/slipneff/tg-spam/internal/models"
	"gorm.io/gorm"
)

func (s *Storage) CreateChannel(ctx context.Context, input *models.Channel) (*models.Channel, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.Create(input).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			err := tr.Model(&models.Channel{}).Where("name = ?", input.Name).First(&input).Error
			if err != nil {
				return nil, err
			}
		}
		return nil, err
	}

	return input, nil
}

func (s *Storage) GetChannels(ctx context.Context) ([]*models.Channel, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	var channels []*models.Channel
	err := tr.Find(&channels).Error
	if err != nil {
		return nil, err
	}

	return channels, nil
}

func (s *Storage) SetLastMessageID(ctx context.Context, name string, id int) error {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.Model(&models.Channel{}).Where("name = ?", name).Update("last_message_id", id).Error
	if err != nil {
		return err
	}
	return nil
}
