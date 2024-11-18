package bot

import (
	"github.com/mymmrac/telego"
	"github.com/slipneff/tg-spam/internal/service"
)

type Bot struct {
	*telego.Bot
	service *service.Service
	Chat    map[int64]TgScene
}

type TgScene int

const (
	SceneMain TgScene = iota
	SceneAuth
)

func New(bot *telego.Bot, service *service.Service) *Bot {
	return &Bot{Bot: bot, service: service, Chat: make(map[int64]TgScene)}
}
