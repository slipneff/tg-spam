package main

import (
	"context"

	"github.com/slipneff/gogger"
	"github.com/slipneff/tg-spam/internal/di"
	"github.com/slipneff/tg-spam/internal/utils/config"
	"github.com/slipneff/tg-spam/internal/utils/flags"
)

func main() {
	flags := flags.MustParseFlags()
	cfg := config.MustLoadConfig(flags.EnvMode, flags.ConfigPath)
	gogger.ConfigureZeroLogger()
	ctx := context.Background()
	container := di.New(ctx, cfg)

	err := container.GetService().CatchingChannels(ctx)
	if err != nil {
		panic(err)
	}
}
