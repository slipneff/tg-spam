package tg

import (
	"github.com/slipneff/tg-spam/internal/utils/config"
)

type Client struct {
	Name string
	cfg  *config.Config
}

func NewClient(name string, cfg *config.Config) *Client {
	return &Client{Name: name, cfg: cfg}
}
