package sql

import (
	"context"

	"github.com/google/uuid"
	"github.com/slipneff/tg-spam/internal/models"
)

func (s *Storage) CreateAccount(ctx context.Context, input *models.Account) (*models.Account, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.Create(input).Error
	if err != nil {
		return nil, err
	}

	return input, nil
}

func (s *Storage) BatchCreateAccounts(ctx context.Context, accounts []*models.Account) error {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	return tr.Create(accounts).Error
}

func (s *Storage) GetAccount(ctx context.Context, id uuid.UUID) (*models.Account, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	var account models.Account
	err := tr.First(&account, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *Storage) GetAccountByEmail(ctx context.Context, email string) (*models.Account, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	var account models.Account
	err := tr.First(&account, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *Storage) GetAccounts(ctx context.Context) ([]*models.Account, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	var accounts []*models.Account
	err := tr.Find(&accounts).Error
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *Storage) GetAccountsWithToken(ctx context.Context) ([]*models.Account, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	var accounts []*models.Account
	err := tr.Find(&accounts, "auth = true").Error
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *Storage) GetAccountsWithoutSecret(ctx context.Context) ([]*models.Account, error) {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	var accounts []*models.Account
	err := tr.Find(&accounts, "client_secret = false").Error
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *Storage) ConfirmAuth(ctx context.Context, id uuid.UUID) error {
	tr := s.getter.DefaultTrOrDB(ctx, s.db).WithContext(ctx)
	err := tr.Model(&models.Account{}).Where("id = ?", id).Update("auth", true).Error
	if err != nil {
		return err
	}
	return nil
}
