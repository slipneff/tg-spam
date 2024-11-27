package sql

import (
	"fmt"

	"github.com/slipneff/tg-spam/internal/models"
	"github.com/slipneff/tg-spam/internal/utils/config"

	trmgorm "github.com/avito-tech/go-transaction-manager/gorm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db     *gorm.DB
	getter *trmgorm.CtxGetter
}

func New(db *gorm.DB, getter *trmgorm.CtxGetter) *Storage {
	return &Storage{
		db:     db,
		getter: getter,
	}
}

func buildDSN(cfg *config.Config) string {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.Name, cfg.DB.SSLMode,
	)

	return dsn
}

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(buildDSN(cfg)), &gorm.Config{
		TranslateError: true,
	})
}

func MustNewPostgresDB(cfg *config.Config) *gorm.DB {
	db, err := NewPostgresDB(cfg)
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.Channel{}, &models.Session{})
	if err != nil {
		panic(err)
	}
	return db
}
