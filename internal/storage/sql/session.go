package sql

import (
	"context"

	"github.com/slipneff/tg-spam/internal/models"
)

func (s *Storage) GetSessions(ctx context.Context, n int) ([]*models.Session, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	var sessions []*models.Session
	err := tr.Find(&sessions).Limit(n).Error
	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (s *Storage) GetSessionById(ctx context.Context, id string) (*models.Session, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	var session models.Session
	err := tr.Where("id = ?", id).First(&session).Error
	if err != nil {
		return nil, err
	}

	return &session, nil
}
