package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	Email        string    `gorm:"column:email;unique"`
	Password     string    `gorm:"column:password"`
	Recovery     string    `gorm:"column:recovery"`
	DeveloperKey string    `gorm:"column:developer_key"`
	Auth         bool      `gorm:"column:auth"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		a.ID = id
	}

	return nil
}

func (a *Account) String() string {
	b := ""
	if a.Auth {
		b = "READY"
	} else {
		b = "NO AUTH"
	}

	return a.Email + " - " + b
}
