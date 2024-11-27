package di

import (
	"context"

	trmgorm "github.com/avito-tech/go-transaction-manager/gorm"
	"github.com/avito-tech/go-transaction-manager/trm"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	"github.com/gotd/td/telegram"

	"github.com/slipneff/tg-spam/internal/service"
	"github.com/slipneff/tg-spam/internal/storage/sql"
	"github.com/slipneff/tg-spam/internal/worker"

	"gorm.io/gorm"

	"github.com/slipneff/tg-spam/internal/utils/config"
)

type Container struct {
	cfg *config.Config
	ctx context.Context

	service            *service.Service
	worker             *worker.Worker
	storage            *sql.Storage
	db                 *gorm.DB
	transactionManager trm.Manager
}

func New(ctx context.Context, cfg *config.Config) *Container {
	return &Container{cfg: cfg, ctx: ctx}
}

func (c *Container) GetPostgresDB() *sql.Storage {
	return get(&c.storage, func() *sql.Storage {
		return sql.New(c.GetDB(), trmgorm.DefaultCtxGetter)
	})
}

func (c *Container) GetDB() *gorm.DB {
	return get(&c.db, func() *gorm.DB {
		return sql.MustNewPostgresDB(c.cfg)
	})
}

func (c *Container) GetTransactionManager() trm.Manager {
	return get(&c.transactionManager, func() trm.Manager {
		return manager.Must(trmgorm.NewDefaultFactory(c.db))
	})
}

func (c *Container) GetService() *service.Service {

	return get(&c.service, func() *service.Service {
		return service.NewService(c.GetWorker(), c.cfg)
	})
}
func (c *Container) GetWorker() *worker.Worker {
	client := telegram.NewClient(c.cfg.Telegram.AppID, c.cfg.Telegram.AppHash, telegram.Options{})
	return get(&c.worker, func() *worker.Worker {
		return worker.NewWorker(client, c.cfg)
	})
}

func get[T comparable](obj *T, builder func() T) T {
	if *obj != *new(T) {
		return *obj
	}

	*obj = builder()
	return *obj
}
