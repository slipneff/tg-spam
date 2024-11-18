package main

import (
	"context"
	"fmt"

	"github.com/slipneff/gogger"
	"github.com/slipneff/gogger/log"
	"github.com/slipneff/tg-spam/internal/di"
	"github.com/slipneff/tg-spam/internal/utils/config"
	"github.com/slipneff/tg-spam/internal/utils/flags"
)

func main() {
	flags := flags.MustParseFlags()
	cfg := config.MustLoadConfig(flags.EnvMode, flags.ConfigPath)
	gogger.ConfigureZeroLogger()

	container := di.New(context.Background(), cfg)
	log.Info(fmt.Sprintf("Server starting at %s:%d", cfg.Host, cfg.Port))

	bot := container.NewBot()

	bot.Auth(context.Background())
}
